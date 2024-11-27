package runner

import (
	"context"
	"time"

	"github.com/iwashi623/kinben/options"
)

const DefaultTimeout = 300 * time.Second

type BenchRunner interface {
	IsuconName() string
	Run(ctx context.Context, opt *options.BenchOption) (string, error)
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

func Run(ctx context.Context, opt *options.BenchOption) (string, error) {
	return br.Run(ctx, opt)
}
