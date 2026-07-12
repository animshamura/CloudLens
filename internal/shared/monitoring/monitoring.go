package monitoring

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

// Handler exposes a basic Prometheus-compatible metrics endpoint.
type Handler struct {
	requests atomic.Uint64
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.requests.Add(1)
	w.Header().Set("Content-Type", "text/plain; version=0.0.4")
	_, _ = fmt.Fprintf(w, "# HELP http_requests_total Total number of HTTP requests served\n")
	_, _ = fmt.Fprintf(w, "# TYPE http_requests_total counter\n")
	_, _ = fmt.Fprintf(w, "http_requests_total %d\n", h.requests.Load())
}
