package kayaclisten80

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"sync"

	"github.com/iwashi623/kinben/options"
	"github.com/iwashi623/kinben/runner"
)

const (
	IsuconName = "kayac-listen80"
)

type listen80BenchRunner struct {
}

func NewBenchRunner() *listen80BenchRunner {
	return &listen80BenchRunner{}
}

func (bm *listen80BenchRunner) IsuconName() string {
	return IsuconName
}

func (bm *listen80BenchRunner) Run(ctx context.Context, opt *options.BenchOption) (*runner.BenchResult, error) {
	cmd := exec.CommandContext(ctx, "./bench", "-target-url", opt.GetTargetHost())

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to get stdout pipe: %w", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to get stderr pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start command: %w", err)
	}

	var output, errors []byte
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		out, _ := io.ReadAll(stdout)
		output = append(output, out...)
	}()

	go func() {
		defer wg.Done()
		errOut, _ := io.ReadAll(stderr)
		errors = append(errors, errOut...)
	}()

	wg.Wait()

	if err := cmd.Wait(); err != nil {
		if os.IsPermission(err) {
			return nil, fmt.Errorf("failed to execute command: %w", err)
		}

		log.Println(string(errors))
	}

	result, err := bm.parseBenchResult(string(output), opt.GetTargetHost())
	if err != nil {
		return nil, fmt.Errorf("failed to parse bench result: %w", err)
	}

	return result, nil
}

func (bm *listen80BenchRunner) parseBenchResult(logOutput, target string) (*runner.BenchResult, error) {
	scoreRegex := regexp.MustCompile(`SCORE:\s*(-?\d+)`)
	scoreMatch := scoreRegex.FindStringSubmatch(logOutput)
	if scoreMatch == nil {
		return nil, fmt.Errorf("failed to parse score")
	}
	score, err := strconv.Atoi(scoreMatch[1])
	if err != nil {
		return nil, fmt.Errorf("invalid score value: %w", err)
	}

	resultRegex := regexp.MustCompile(`RESULT:\s*(.*)`)
	resultMatch := resultRegex.FindStringSubmatch(logOutput)
	if resultMatch == nil {
		return nil, fmt.Errorf("failed to parse result")
	}
	result := resultMatch[1]

	// BenchResultを作成
	return &runner.BenchResult{
		IsuconName: IsuconName,
		Target:     target,
		Score:      score,
		Result:     result,
		Output:     logOutput,
	}, nil
}
