package runner

import (
	"context"
	"time"

	"github.com/iwashi623/kinben/options"
)

const DefaultTimeout = 300 * time.Second

type BenchRunner interface {
	IsuconName() string
	Run(ctx context.Context, opt *options.BenchOption) (*BenchResult, error)
}

var br BenchRunner

type BenchCreateFunc func() (BenchRunner, error)

func RegisterBenchRunner(f BenchCreateFunc) (err error) {
	br, err = f()
	if err != nil {
		return err
	}
	return nil
}

func Run(ctx context.Context, opt *options.BenchOption) (*BenchResult, error) {
	return br.Run(ctx, opt)
}

type BenchResult struct {
	IsuconName string `json:"isucon_name"`
	Target     string `json:"target"`
	Score      int    `json:"score"`
	Result     string `json:"result"`
	Output     string `json:"output"`
}
