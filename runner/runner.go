package runner

import (
	"context"

	"github.com/iwashi623/kinben/options"
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
