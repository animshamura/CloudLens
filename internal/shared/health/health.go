package health

import "net/http"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

// ServeHTTP handles health check requests
// @Summary Health check endpoint
// @Description Returns liveness and readiness status
// @Produce plain
// @Success 200 {string} string "ok"
// @Success 200 {string} string "ready"
// @Failure 404 {string} string "Not Found"
// @Router /health/live [get]
// @Router /health/ready [get]
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/health/live":
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	case "/health/ready":
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ready"))
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
