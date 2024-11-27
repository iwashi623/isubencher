package kayaclisten80

import (
	"net/http"

	"github.com/iwashi623/kinben/bench"
	"github.com/iwashi623/kinben/options"
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

	out, err := bench.Run(opt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(out))
}
