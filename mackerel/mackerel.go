package mackerel

import (
	"context"
	"fmt"
	"time"

	"github.com/iwashi623/kinben/exporter"

	"github.com/mackerelio/mackerel-client-go"
)

type MackerelExporter struct {
	ServiceName string
	Client      *mackerel.Client
}

var _ exporter.Exporter = (*MackerelExporter)(nil)

func NewMackerelClient(apiKey string) *mackerel.Client {
	return mackerel.NewClient(apiKey)
}

func NewMackerelExporter(
	serviceName string,
	client *mackerel.Client,
) *MackerelExporter {

	return &MackerelExporter{
		ServiceName: serviceName,
		Client:      client,
	}
}

func (m *MackerelExporter) Export(ctx context.Context, params exporter.ExportParams) error {
	fmt.Println(params)
	err := m.Client.PostServiceMetricValues(m.ServiceName, []*mackerel.MetricValue{
		{
			Name:  params.TeamName,
			Value: params.Score,
			Time:  time.Now().Unix(),
		},
	})

	return err
}

func (m *MackerelExporter) GetExporterName() string {
	return "mackerel"
}
