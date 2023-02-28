package main

type Metric struct {
	Name   string            `json:"name"`
	Labels map[string]string `json:"labels"`
	Value  float64           `json:"value"`
}

type MonitoringStruct struct {
	Metrics []Metric `json:"metrics"`
}
