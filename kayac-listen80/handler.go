package kayaclisten80

import (
	"context"
	"errors"
	"net/http"

	"github.com/iwashi623/kinben/options"
	"github.com/iwashi623/kinben/runner"
)

type listen80Hander struct {
}

func NewHandler() *listen80Hander {
	return &listen80Hander{}
}

func (h *listen80Hander) Handle(ctx context.Context, req *http.Request) ([]byte, error) {
	targetHost := req.URL.Query().Get("target-host")
	if targetHost == "" {
		return nil, errors.New("target-host is required")
	}

	opt := options.NewBenchOption(
		targetHost,
	)

	result, err := runner.Run(ctx, opt)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, errors.New("request timed out")
		}
		return nil, err
	}

	json, err := result.ToJSON()
	if err != nil {
		return nil, err
	}

	return json, nil
}
