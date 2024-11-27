package kayaclisten80

import (
	"context"
	"net/http"

	"github.com/iwashi623/kinben/options"
	"github.com/iwashi623/kinben/runner"
)

func BenchHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), runner.DefaultTimeout)
	defer cancel()

	targetHost := r.URL.Query().Get("target-ip")
	if targetHost == "" {
		http.Error(w, "target-ip is required", http.StatusBadRequest)
		return
	}

	opt := options.NewBenchOption(
		targetHost,
	)

	out, err := runner.Run(ctx, opt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(out))
}
