package kinben

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	kayaclisten80 "github.com/iwashi623/kinben/kayac-listen80"
	"github.com/iwashi623/kinben/options"
	"github.com/iwashi623/kinben/runner"
)

type BenchHandler interface {
	Handle(req *http.Request) error
}

type benchDefinition struct {
	benchRunner runner.RunnerCreateFunc
	handler     BenchHandler
}

var definitions = map[string]*benchDefinition{
	kayaclisten80.IsuconName: {
		benchRunner: WrapKayaclisten80NewBenchRunner,
		handler:     WrapKayaclisten80NewBenchHandler(),
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
	mux.HandleFunc("/bench", k.createBenchHandler(h))
	return nil
}

func newHandler(name string) (BenchHandler, error) {
	if bd, exists := definitions[name]; exists {
		return bd.handler, nil
	}
	return nil, fmt.Errorf("no competition")
}

func (k *kinben) createBenchHandler(h BenchHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
}

func WrapKayaclisten80NewBenchRunner() (runner.Runner, error) {
	return kayaclisten80.NewBenchRunner()
}

func WrapKayaclisten80NewBenchHandler() BenchHandler {
	return kayaclisten80.NewHandler()
}
