package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sstallion/go-hid"
)

var (
	debug       bool
	showVersion bool
	port        string
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var (
	co2 = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "co2",
		Help: "Air CO2 counter",
	})
	temperature = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "temp",
		Help: "Air Temperature in celsius degrees",
	})
)

func main() {
	// command line arg
	flag.StringVar(&port, "port", "2112", "Port number to start http server")
	flag.BoolVar(&debug, "debug", false, "Enable debug output")
	flag.BoolVar(&showVersion, "version", false, "Show version")
	flag.Parse()

	// show version and exit
	if showVersion {
		fmt.Printf("%s verion %s, commit %s, built at %s", path.Base(os.Args[0]), version, commit, date)
		os.Exit(0)
	}

	// start web server
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			log.Fatalf("could not start http server: %v", err)
		}
	}()

	// init USB HID
	if err := hid.Init(); err != nil {
		log.Fatal(err)
	}

	// Open device
	d, err := hid.OpenFirst(0x04d9, 0xa052)
	if err != nil {
		fmt.Printf("can't open device - %v\n", err)
		os.Exit(1)
	}

	// Print USB HID debug info
	if debug {
		DeviceInfo(d)
	}

	magic := make([]byte, 8)
	_, err = d.SendFeatureReport(magic)
	if err != nil {
		log.Fatal(err)
	}

	// Loop to get values in prometheus counters
	for {
		getValues(d)
	}
}

func DeviceInfo(d *hid.Device) {
	// Read the Manufacturer String.
	s, err := d.GetMfrStr()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Manufacturer String: %s\n", s)

	// Read the Product String.
	s, err = d.GetProductStr()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Product String: %s\n", s)

	// Read the Serial Number String.
	s, err = d.GetSerialNbr()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Serial Number String: %s\n", s)
}

func getValues(dev *hid.Device) {
	buf := make([]byte, 8)

	i, err := dev.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	if i != len(buf) {
		log.Fatal("wrong read bug size")
	}

	decode(buf)
}

func decode(buf []byte) {
	metric, value := buf[0], buf[1]
	valueThousands, _ := buf[2], buf[3]

	if buf[4] != 0x0d {
		log.Println("Unexpected data from device")
		return
	}

	metricName := ""
	var valueShifted int32
	valueShifted = int32(value) << 8
	metricValue := valueShifted + int32(valueThousands)
	var metricValueF float64
	switch metric {
	case 0x42:
		metricName = "Temperature"
		metricValueF = float64(metricValue)/16 - 273.15
		temperature.Set(metricValueF)
	case 0x50:
		metricName = "CO2"
		metricValueF = float64(metricValue)
		co2.Set(metricValueF)
	case 0x44:
		metricName = "Humidity"
		metricValueF = float64(metricValue)
	default:
		metricName = "Unknown"
		return
	}

	if debug {
		log.Printf("0x%X %s (%.1f)\n", metric, metricName, metricValueF)
	}
}
