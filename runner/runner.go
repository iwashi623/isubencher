package runner

import (
	"context"
	"fmt"
	"regexp"

	"github.com/iwashi623/kinben/exporter"
	"github.com/iwashi623/kinben/options"
	"github.com/iwashi623/kinben/teamsheet"
)

type Runner interface {
	IsuconName() string
	Run(ctx context.Context, opt *options.BenchOption) (*BenchResult, error)
}

type BenchResult struct {
	IsuconName string
	Target     string
	Score      int
	Result     string
	Output     string
}

type runner struct {
	runner   Runner
	sheet    teamsheet.TeamSheet
	exporter exporter.Exporter
}

func NewRunner(
	r Runner,
	s teamsheet.TeamSheet,
	e exporter.Exporter,
) *runner {
	return &runner{
		runner:   r,
		sheet:    s,
		exporter: e,
	}
}

func (r *runner) Run(ctx context.Context, opt *options.BenchOption) (*BenchResult, error) {
	hostIP, err := extractIPAddress(opt.GetTargetHost())
	if err != nil {
		return nil, fmt.Errorf("failed to extract IP address: %w", err)
	}

	teamName, err := r.sheet.GetTeamNameByIP(hostIP)
	if err != nil {
		return nil, fmt.Errorf("failed to get team name: %w", err)
	}

	fmt.Printf("team name: %s\n", teamName)
	result, err := r.runner.Run(ctx, opt)
	if err != nil {
		return nil, err
	}

	if teamName == "" {
		return nil, fmt.Errorf("no team name found for IP: %s", hostIP)
	}

	if err := r.exporter.Export(exporter.ExportParams{
		TeamName: teamName,
		Score:    result.Score,
	}); err != nil {
		return nil, fmt.Errorf("failed to export: %w", err)
	}

	return result, nil
}

func (r *runner) IsuconName() string {
	return r.runner.IsuconName()
}

func extractIPAddress(input string) (string, error) {
	re := regexp.MustCompile(`https?://(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`)

	match := re.FindStringSubmatch(input)
	if len(match) < 2 {
		return "", fmt.Errorf("no IP address found in input")
	}

	return match[1], nil
}
