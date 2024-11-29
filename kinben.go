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

type BenchHandler interface {
	Handle(ctx context.Context, req *http.Request) ([]byte, error)
}

type benchDefinition struct {
	benchRunner runner.Runner
	handler     BenchHandler
}

var definitions = map[string]*benchDefinition{
	kayaclisten80.IsuconName: {
		benchRunner: WrapKayaclisten80NewBenchRunner(),
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
	k := &kinben{
		s: &http.Server{
			Addr: ":" + port,
		},
		isuconName: isuconName,
	}

	mux := http.NewServeMux()
	if err := k.registerRoutes(mux); err != nil {
		return nil, err
	}
	k.s.Handler = mux

	err := RegisterBenchRunner(isuconName)
	if err != nil {
		return nil, err
	}

	return k, nil
}

func RegisterBenchRunner(isuconName string) error {
	if bd, exists := definitions[isuconName]; exists {
		return runner.RegisterBenchRunner(bd.benchRunner)
	}
	return fmt.Errorf("no competition")
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

		result, err := h.Handle(ctx, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// resultをjsonで返す
		if _, err := w.Write(result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (k *kinben) StartServer() error {
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

func WrapKayaclisten80NewBenchRunner() runner.Runner {
	return kayaclisten80.NewBenchRunner()
}

func WrapKayaclisten80NewBenchHandler() BenchHandler {
	return kayaclisten80.NewHandler()
}
