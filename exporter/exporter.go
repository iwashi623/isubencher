package exporter

type Exporter interface {
	Export(params ExportParams) error
	GetExporterName() string
}

type ExportParams struct {
	TeamName string
	Score    int
}

type exporter struct {
}

func NewExporter() *exporter {
	return &exporter{}
}

func (e *exporter) Export(params ExportParams) error {
	return nil
}

func (e *exporter) GetExporterName() string {
	return "nil exporter"
}

type ExporterCreateFunc func() (Exporter, error)
