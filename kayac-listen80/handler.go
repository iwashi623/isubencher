package kayaclisten80

import (
	"net/http"
	"strconv"

	"github.com/iwashi623/kinben/bench"
	"github.com/iwashi623/kinben/options"
)

func BenchHandler(w http.ResponseWriter, r *http.Request) {
	targetHost := r.URL.Query().Get("target-host")
	if targetHost == "" {
		http.Error(w, "target-ip is required", http.StatusBadRequest)
		return
	}
	sslEnabled := r.URL.Query().Get("ssl-enabled")
	sslFlag, err := strconv.ParseBool(sslEnabled)
	if err != nil {
		http.Error(w, "ssl-enabled is invalid", http.StatusBadRequest)
		return
	}

	benchProtcol := "http"
	if sslFlag {
		benchProtcol = "https"
	}

	opt := options.NewBenchOption(
		targetHost,
		benchProtcol,
	)

	out, err := bench.Run(opt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(out))
}
