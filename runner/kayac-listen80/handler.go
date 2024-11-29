package kayaclisten80

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/iwashi623/kinben/exporter"
	"github.com/iwashi623/kinben/options"
	"github.com/iwashi623/kinben/response"
	kinbenrunner "github.com/iwashi623/kinben/runner"
	"github.com/iwashi623/kinben/teamsheet"
)

type listen80Hander struct {
	runner   kinbenrunner.Runner
	sheet    teamsheet.TeamSheet
	exporter exporter.Exporter
}

func NewHandler(
	runner kinbenrunner.Runner,
	sheet teamsheet.TeamSheet,
	exporter exporter.Exporter,
) *listen80Hander {
	return &listen80Hander{
		runner:   runner,
		sheet:    sheet,
		exporter: exporter,
	}
}

func (h *listen80Hander) Handle(
	ctx context.Context,
	req *http.Request,
) (*response.BenchResponse, error) {
	targetHost := req.URL.Query().Get("target-host")
	if targetHost == "" {
		return nil, errors.New("target-host is required")
	}

	opt := options.NewBenchOption(
		targetHost,
	)

	hostIP, err := extractIPAddress(targetHost)
	if err != nil {
		return nil, fmt.Errorf("failed to extract IP address: %w", err)
	}

	teamName := h.sheet.GetTeamNameByIP(hostIP)

	result, err := h.runner.Run(ctx, opt)
	if err != nil {
		return nil, err
	}

	if teamName == "" {
		return nil, fmt.Errorf("no team name found for IP: %s", hostIP)
	}

	if err := h.exporter.Export(exporter.ExportParams{
		TeamName: teamName,
		Score:    result.Score,
	}); err != nil {
		return nil, fmt.Errorf("failed to export: %w", err)
	}

	return response.NewBenchResponse(result), nil
}

func extractIPAddress(input string) (string, error) {
	re := regexp.MustCompile(`https?://(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`)

	match := re.FindStringSubmatch(input)
	if len(match) < 2 {
		return "", fmt.Errorf("no IP address found in input")
	}

	return match[1], nil
}
