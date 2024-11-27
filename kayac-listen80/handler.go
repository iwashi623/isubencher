package kayaclisten80

import (
	"net/http"

	"github.com/iwashi623/kinben/options"
	"github.com/iwashi623/kinben/runner"
)

func BenchHandler(w http.ResponseWriter, r *http.Request) {
	targetHost := r.URL.Query().Get("target-host")
	if targetHost == "" {
		http.Error(w, "target-ip is required", http.StatusBadRequest)
		return
	}

	opt := options.NewBenchOption(
		targetHost,
	)

	out, err := runner.Run(opt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(out))
}
