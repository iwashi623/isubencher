package kayaclisten80

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/iwashi623/kinben/options"
	"github.com/iwashi623/kinben/runner"
)

func BenchHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), runner.DefaultTimeout)
	defer cancel()

	targetHost := r.URL.Query().Get("target-host")
	if targetHost == "" {
		http.Error(w, "target-host is required", http.StatusBadRequest)
		return
	}

	opt := options.NewBenchOption(
		targetHost,
	)

	result, err := runner.Run(ctx, opt)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			http.Error(w, "Request timed out", http.StatusGatewayTimeout)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// resultをjsonで返す
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
