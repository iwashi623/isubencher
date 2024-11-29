package runner

import (
	"context"
	"time"

	"github.com/iwashi623/kinben/options"
	"github.com/iwashi623/kinben/result"
)

const DefaultTimeout = 300 * time.Second

type Runner interface {
	IsuconName() string
	Run(ctx context.Context, opt *options.BenchOption) (*result.BenchResult, error)
}

var br Runner

type RunnerCreateFunc func() (Runner, error)

func RegisterBenchRunner(f RunnerCreateFunc) (err error) {
	br, err = f()
	if err != nil {
		return err
	}
	return nil
}

func Run(ctx context.Context, opt *options.BenchOption) (*result.BenchResult, error) {
	return br.Run(ctx, opt)
}
