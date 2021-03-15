package influxdbv2

import (
	"errors"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/loadimpact/k6/output"
	"github.com/loadimpact/k6/stats"
	"os"
	"time"
)

func init() {
	output.RegisterExtension("influxdbv2", func(params output.Params) (output.Output, error) {
		return &Output{params: params}, nil
	})
}

type Output struct {
	params   output.Params
	client   influxdb2.Client
	writeAPI api.WriteAPI
}

func (o *Output) Description() string {
	return "Output to influxDB(V2)."
}

func (o *Output) Start() error {
	influxDBv2Url := os.Getenv("influxDBv2Url")
	influxDBv2Token := os.Getenv("influxDBv2Token")
	influxDBv2Organization := os.Getenv("influxDBv2Organization")
	influxDBv2Bucket := os.Getenv("influxDBv2Bucket")
	if influxDBv2Url == "" || influxDBv2Token == "" ||
		influxDBv2Organization == "" || influxDBv2Bucket == "" {
		return errors.New(" Dependent environment variables are not specified, " +
			"please check if the environment variable exists.")
	}
	o.client = influxdb2.NewClientWithOptions(influxDBv2Url, influxDBv2Token, influxdb2.DefaultOptions().
		SetHTTPRequestTimeout(60))
	o.writeAPI = o.client.WriteAPI(influxDBv2Organization, influxDBv2Bucket)
	return nil
}

func (o *Output) AddMetricSamples(sampleContainers []stats.SampleContainer) {
	for _, sc := range sampleContainers {
		for _, sample := range sc.GetSamples() {
			point := Sample2Point(sample)
			o.writeAPI.WritePoint(point)
		}
	}
}

func Sample2Point(sample stats.Sample) *write.Point {
	tags := map[string]string{}
	fields := map[string]interface{}{}
	for tagK, tagV := range sample.Tags.CloneTags() {
		fields[tagK] = tagV
	}
	fields["_value"] = sample.Value
	fields["debug"] = sample.Metric.Name
	if sample.Metric.Name == "http_req_duration" {
		fields["_finished"] = sample.Time.Add(time.Duration(sample.Value * 1_000_000))
	}
	typeBytes, _ := sample.Metric.Type.MarshalText()
	tags["type"] = string(typeBytes)
	point := influxdb2.NewPoint(sample.Metric.Name, tags, fields, sample.Time)
	return point
}

func (o *Output) Stop() error {
	o.client.Close()
	return nil
}
