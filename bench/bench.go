package bench

import (
	"github.com/iwashi623/kinben/options"
)

type BenchMarker interface {
	IsuconName() string
	Run(opt *options.BenchOption) (string, error)
}

var bench BenchMarker

type BenchCreateFunc func() (BenchMarker, error)

func RegisterBenchMarker(f BenchCreateFunc) (err error) {
	bench, err = f()
	if err != nil {
		return err
	}

	return nil
}

func Run(opt *options.BenchOption) (string, error) {
	return bench.Run(opt)
}
