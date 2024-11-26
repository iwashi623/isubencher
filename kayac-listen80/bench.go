package kayaclisten80

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/iwashi623/isubencher/options"
)

type BenchMarker struct {
}

func NewBenchMarker() (*BenchMarker, error) {
	return &BenchMarker{}, nil
}

func (bm *BenchMarker) Run(opt *options.BenchOption) (string, error) {
	cmd := exec.Command("./bench", "-target-url", opt.GetBenchProtcol()+"://"+opt.GetTargetHost())

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
