package gateway

import (
	"net/http"

	_ "github.com/cloudlens/cloud-native-ecommerce/docs"
	"github.com/cloudlens/cloud-native-ecommerce/internal/shared/health"
	"github.com/cloudlens/cloud-native-ecommerce/internal/shared/monitoring"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	h := health.NewHandler()
	
	// Health check routes
	mux.Handle("/health/", h)
	
	// Root endpoint
	mux.HandleFunc("/", handleRoot)

	// Monitoring endpoint
	metricsHandler := monitoring.NewHandler()
	mux.Handle("/metrics", metricsHandler)
	
	// Swagger UI with embedded documentation
	mux.Handle("/swagger/", httpSwagger.Handler())

	return WithTracing(mux)
}

// handleRoot godoc
// @Summary Welcome to API Gateway
// @Description Returns a welcome message
// @Produce plain
// @Success 200 {string} string "cloud-native-ecommerce api gateway"
// @Router / [get]
func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("cloud-native-ecommerce api gateway"))
}
