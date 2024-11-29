package mackerel

import "github.com/iwashi623/kinben/exporter"

type Mackerel struct {
}

func NewMackerel() *Mackerel {
	return &Mackerel{}
}

func (m *Mackerel) Export(params exporter.ExportParams) error {
	return nil
}
