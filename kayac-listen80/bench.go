package kayaclisten80

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/iwashi623/kinben/options"
	"github.com/iwashi623/kinben/runner"
)

const (
	IsuconName = "kayac-listen80"
)

type listen80BenchRunner struct {
}

func NewBenchRunner() (runner.BenchRunner, error) {
	return &listen80BenchRunner{}, nil
}

func (bm *listen80BenchRunner) IsuconName() string {
	return IsuconName
}

func (bm *listen80BenchRunner) Run(ctx context.Context, opt *options.BenchOption) (string, error) {
	cmd := exec.CommandContext(ctx, "./bench", "-target-url", opt.GetTargetHost())

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		if os.IsPermission(err) {
			fmt.Println("実行権限が不足しています。権限を確認してください。")
		}
		return "", err
	}

	return string(out), nil
}
