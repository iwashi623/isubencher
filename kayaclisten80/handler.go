package kayaclisten80

import (
	"context"
	"errors"
	"net/http"

	"github.com/iwashi623/kinben/options"
	"github.com/iwashi623/kinben/response"
	kinbenrunner "github.com/iwashi623/kinben/runner"
)

type listen80Hander struct {
	runner kinbenrunner.Runner
}

func NewHandler(
	runner kinbenrunner.Runner,
) *listen80Hander {
	return &listen80Hander{
		runner: runner,
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

	result, err := h.runner.Run(ctx, opt)
	if err != nil {
		return nil, err
	}

	return response.NewBenchResponse(result), nil
}
