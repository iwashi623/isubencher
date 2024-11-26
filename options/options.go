package options

type BenchOption struct {
	// IP address of the target server
	targetHost string
	// Whether to use SSL
	sslEnabled bool
	// http or https
	benchProtcol string
}

func NewBenchOption(
	targetHost string,
	benchProtcol string,
	opt ...OptionFunc,
) *BenchOption {
	option := &BenchOption{
		targetHost:   targetHost,
		benchProtcol: benchProtcol,
	}

	for _, o := range opt {
		o(option)
	}

	return option
}

type OptionFunc func(*BenchOption)

func WithTargetHost(targetHost string) OptionFunc {
	return func(opt *BenchOption) {
		opt.targetHost = targetHost
	}
}

func (o *BenchOption) GetTargetHost() string {
	return o.targetHost
}

func WithSslEnabled(sslEnabled bool) OptionFunc {
	return func(opt *BenchOption) {
		opt.sslEnabled = sslEnabled
	}
}

func WithBenchProtcol(benchProtcol string) OptionFunc {
	return func(opt *BenchOption) {
		opt.benchProtcol = benchProtcol
	}
}

func (o *BenchOption) GetBenchProtcol() string {
	return o.benchProtcol
}

type BenchMarker interface {
	Run(opt ...OptionFunc) (string, error)
}
