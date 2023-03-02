package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	ycsdk "github.com/yandex-cloud/go-sdk"
	"io"
	"log"
	"net/http"
	"time"
)

var ServiceAccountKeyFile string
var Token string

var folderID string

func main() {

	flag.StringVar(&ServiceAccountKeyFile, "sa-key-file", "", "Service account key file")
	flag.StringVar(&Token, "token", "", "Yandex.Cloud token")
	flag.StringVar(&folderID, "folder-id", "", "Yandex.Cloud folder ID")
	flag.Parse()

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
	iamToken, err := tokenMiddleware.GetIAMToken(ctx, false)
	if err != nil {
		log.Fatal(err)
	}
	sendMetrics(ctx, iamToken)

	ticker := time.NewTicker(15 * time.Second)
	quit := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			iamToken, err := tokenMiddleware.GetIAMToken(ctx, false)
			if err != nil {
				log.Fatal(err)
			}
			sendMetrics(ctx, iamToken)
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

var endpoint = "monitoring.api.cloud.yandex.net"

// sendMetrics sends metrics to Yandex.Cloud Monitoring with forming http query params and marshaling
// metrics to json

//'https://monitoring.api.cloud.yandex.net/monitoring/v2/data/write?folderId=aoe6vrq0g3svvs3uf62u&service=custom' > output.json

func sendMetrics(ctx context.Context, iamToken string) {
	data := MonitoringStruct{}
	data1 := Metric{
		Name:   "Temp",
		Labels: map[string]string{},
		Value:  10,
	}
	data.Metrics = append(data.Metrics, data1)

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
	//log.Printf("response Status: %s", response.Status)
	//log.Printf("response Body: %s", body)
}
