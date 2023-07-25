package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var (
	ServiceAccountKeyFile string
	Token                 string
	folderID              string
	Source                string
	showVersion           bool
)

const endpoint = "monitoring.api.cloud.yandex.net"
const period = 15 * time.Second

func main() {
	flag.StringVar(&ServiceAccountKeyFile, "sa-key-file", "", "Service account key file")
	flag.StringVar(&Token, "token", "", "Yandex.Cloud token")
	flag.StringVar(&folderID, "folder-id", "", "Yandex.Cloud folder ID")
	flag.StringVar(&Source, "source", "localhost:2112", "co2mon prometheus exporter source")
	flag.BoolVar(&showVersion, "version", false, "Show version")
	flag.Parse()

	// show version and exit
	if showVersion {
		fmt.Printf("%s verion %s, commit %s, built at %s", path.Base(os.Args[0]), version, commit, date)
		os.Exit(0)
	}

	// validate
	if folderID == "" {
		log.Fatal("validate error: folder-id is required")
	}

	if ServiceAccountKeyFile != "" && Token != "" {
		log.Fatal("validate error: only one of 'token' or 'sa-key-file' should be specified")
	}

	// init
	ctx := context.Background()

	credentials, err := Credentials()
	if err != nil {
		log.Fatal(err)
	}

	config := ycsdk.Config{
		Credentials: credentials,
	}
	sdk, err := ycsdk.Build(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	tokenMiddleware := ycsdk.NewIAMTokenMiddleware(sdk, time.Now)

	ticker := time.NewTicker(period)
	quit := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			metrics := getMetrics()
			iamToken, err := tokenMiddleware.GetIAMToken(ctx, false)
			if err != nil {
				log.Fatal(err)
			}
			sendMetrics(ctx, iamToken, metrics)
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func getMetrics() map[string]*dto.MetricFamily {
	resp, err := http.Get("http://" + Source + "/metrics")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var parser expfmt.TextParser
	mf, err := parser.TextToMetricFamilies(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return mf
}

// sendMetrics sends metrics to Yandex.Cloud Monitoring with forming http query params and marshaling
// metrics to json

//'https://monitoring.api.cloud.yandex.net/monitoring/v2/data/write?folderId=aoe6vrq0g3svvs3uf62u&service=custom' > output.json

func sendMetrics(ctx context.Context, iamToken string, metrics map[string]*dto.MetricFamily) {
	// create data
	data := prepareMetrics(metrics)

	// marshal data to json
	buf, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	url := fmt.Sprintf("https://%s/monitoring/v2/data/write", endpoint) //, folderID)

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(buf))
	if err != nil {
		log.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("folderId", folderID)
	q.Add("service", "custom")
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", iamToken))

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != http.StatusOK {
		log.Printf("response Status: %s", response.Status)
		log.Printf("response Body: %s", body)
	}
}

func prepareMetrics(metrics map[string]*dto.MetricFamily) MonitoringStruct {
	result := MonitoringStruct{}
	for _, m := range metrics {
		for _, metric := range m.Metric {
			data1 := Metric{
				Name:   m.GetName(),
				Labels: map[string]string{},
				Value:  metric.GetGauge().GetValue(),
			}
			for _, label := range metric.Label {
				data1.Labels[label.GetName()] = label.GetValue()
			}
			result.Metrics = append(result.Metrics, data1)
		}
	}
	return result
}
