package kinben

import (
	"fmt"
	"net/http"
	"os"

	"github.com/iwashi623/kinben/bench"
	kayaclisten80 "github.com/iwashi623/kinben/kayac-listen80"
)

type Kinben struct {
	isuconName string
	s          http.Server
}

func NewKinben(
	port string,
	isuconName string,
) *Kinben {
	err := RegisterBenchMarker(isuconName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &Kinben{
		s: http.Server{
			Addr: ":" + port,
		},
		isuconName: isuconName,
	}
}

func RegisterBenchMarker(isuconName string) error {
	switch isuconName {
	case "kayac-listen80":
		return bench.RegisterBenchMarker(kayaclisten80.NewBenchMarker)
	default:
		return fmt.Errorf("no competition")
	}
}

func (k *Kinben) StartServer() error {
	h := newHandler(k.isuconName)

	http.HandleFunc("/bench", h)
	if err := k.s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func newHandler(name string) http.HandlerFunc {
	switch name {
	case kayaclisten80.IsuconName:
		return kayaclisten80.BenchHandler
	default:
		return nil
	}
}
