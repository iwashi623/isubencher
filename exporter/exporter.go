package exporter

type Exporter interface {
	Export(params ExportParams) error
}

type ExportParams struct {
	TeamName string
	Score    int
}
