package kinben

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	kayaclisten80 "github.com/iwashi623/kinben/kayac-listen80"
	"github.com/iwashi623/kinben/runner"
)

type benchDefinition struct {
	benchRunner runner.BenchCreateFunc
	handler     http.HandlerFunc
}

var definitions = map[string]*benchDefinition{
	kayaclisten80.IsuconName: {
		benchRunner: kayaclisten80.NewBenchRunner,
		handler:     kayaclisten80.BenchHandler,
	},
}

type kinben struct {
	isuconName string
	s          *http.Server
}

func NewKinben(
	port string,
	isuconName string,
) (*kinben, error) {
	err := RegisterBenchRunner(isuconName)
	if err != nil {
		return nil, err
	}

	return &kinben{
		s: &http.Server{
			Addr: ":" + port,
		},
		isuconName: isuconName,
	}, nil
}

func RegisterBenchRunner(isuconName string) error {
	if bd, exists := definitions[isuconName]; exists {
		return runner.RegisterBenchRunner(bd.benchRunner)
	}
	return fmt.Errorf("no competition")
}

func (k *kinben) StartServer() error {
	mux := http.NewServeMux()
	if err := k.registerRoutes(mux); err != nil {
		return err
	}
	k.s.Handler = mux

	go func() {
		if err := k.s.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return k.s.Shutdown(ctx)
}

func (k *kinben) registerRoutes(mux *http.ServeMux) error {
	h, err := newHandler(k.isuconName)
	if err != nil {
		return fmt.Errorf("failed to create handler: %w", err)
	}
	mux.HandleFunc("/bench", h)
	return nil
}

func newHandler(name string) (http.HandlerFunc, error) {
	if bd, exists := definitions[name]; exists {
		return bd.handler, nil
	}
	return nil, fmt.Errorf("no competition")
}
