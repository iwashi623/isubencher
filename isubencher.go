package isubencher

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	kayaclisten80 "github.com/iwashi623/isubencher/kayac-listen80"
	"github.com/iwashi623/isubencher/options"
)

var bench BenchMarker

type Isubencher struct {
	isuconName string
	s          http.Server
}

func NewIsubencher(
	port string,
	isuconName string,
) *Isubencher {
	bm, err := newBenchMarker(isuconName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bench = bm
	return &Isubencher{
		s: http.Server{
			Addr: ":" + port,
		},
		isuconName: isuconName,
	}
}

func (i *Isubencher) StartServer() error {
	http.HandleFunc("/bench", BenchHandler)
	if err := i.s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

type BenchMarker interface {
	Run(opt *options.BenchOption) (string, error)
}

func newBenchMarker(isuconName string) (BenchMarker, error) {
	switch isuconName {
	case "kayac-listen80":
		return kayaclisten80.NewBenchMarker()
	default:
		return nil, fmt.Errorf("no competition")
	}
}

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
