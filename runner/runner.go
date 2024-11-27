package runner

import (
	"github.com/iwashi623/kinben/options"
)

type BenchRunner interface {
	IsuconName() string
	Run(opt *options.BenchOption) (string, error)
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

func Run(opt *options.BenchOption) (string, error) {
	return br.Run(opt)
}
