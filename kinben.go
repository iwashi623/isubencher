package kinben

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/iwashi623/kinben/exporter/mackerel"
	"github.com/iwashi623/kinben/response"
	kayaclisten80 "github.com/iwashi623/kinben/runner/kayac-listen80"
	"github.com/iwashi623/kinben/teamsheet/spreadsheet"
)

type BenchHandler interface {
	Handle(ctx context.Context, req *http.Request) (*response.BenchResponse, error)
}

const DefaultTimeout = 300 * time.Second

var benchHandlers = map[string]func() BenchHandler{
	kayaclisten80.IsuconName: WrapKayaclisten80NewHandler,
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

	return k, nil
}

func (k *kinben) registerRoutes(mux *http.ServeMux) error {
	h, err := newHandler(k.isuconName)
	if err != nil {
		return fmt.Errorf("failed to create handler: %w", err)
	}

	mux.HandleFunc("/bench", k.createBenchHandlerFunc(h))
	return nil
}

func newHandler(name string) (BenchHandler, error) {
	if h, exists := benchHandlers[name]; exists {
		return h(), nil
	}
	return nil, fmt.Errorf("no competition")
}

func (k *kinben) createBenchHandlerFunc(h BenchHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), DefaultTimeout)
		defer cancel()

		result, err := h.Handle(ctx, r)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				http.Error(w, "request timed out", http.StatusRequestTimeout)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		res, err := result.ToJSON()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// resultをjsonで返す
		if _, err := w.Write(res); err != nil {
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
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return k.s.Shutdown(ctx)
}

func WrapKayaclisten80NewHandler() BenchHandler {
	runner := kayaclisten80.NewBenchRunner()
	sheet := spreadsheet.NewSpreadsheet()
	mackerel := mackerel.NewMackerel()
	return kayaclisten80.NewHandler(runner, sheet, mackerel)
}
