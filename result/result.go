package result

import (
	"bytes"
	"encoding/json"
)

type BenchResult struct {
	IsuconName string `json:"isucon_name"`
	Target     string `json:"target"`
	Score      int    `json:"score"`
	Result     string `json:"result"`
	Output     string `json:"output"`
}

func (bm *BenchResult) ToJSON() ([]byte, error) {
	w := &bytes.Buffer{}
	err := json.NewEncoder(w).Encode(bm)
	if err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}
