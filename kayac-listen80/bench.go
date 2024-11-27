package kayaclisten80

import (
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

type listen80Bench struct {
}

func NewBenchRunner() (runner.BenchRunner, error) {
	return &listen80Bench{}, nil
}

func (bm *listen80Bench) IsuconName() string {
	return IsuconName
}

func (bm *listen80Bench) Run(opt *options.BenchOption) (string, error) {
	cmd := exec.Command("./bench", "-target-url", opt.GetTargetHost())

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
