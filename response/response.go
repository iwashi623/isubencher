package response

import (
	"bytes"
	"encoding/json"

	"github.com/iwashi623/kinben/runner"
)

type BenchResponse struct {
	IsuconName string `json:"isucon_name"`
	Target     string `json:"target"`
	Score      int    `json:"score"`
	Result     string `json:"result"`
	Output     string `json:"output"`
}

func NewBenchResponse(
	result *runner.BenchResult,
) *BenchResponse {
	return &BenchResponse{
		IsuconName: result.IsuconName,
		Target:     result.Target,
		Score:      result.Score,
		Result:     result.Result,
		Output:     result.Output,
	}
}

func (bm *BenchResponse) ToJSON() ([]byte, error) {
	w := &bytes.Buffer{}
	err := json.NewEncoder(w).Encode(bm)
	if err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}
