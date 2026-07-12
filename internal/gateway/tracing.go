package gateway

import (
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type tracingMiddleware struct {
	next http.Handler
}

func (m tracingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer("cloudlens.gateway").Start(r.Context(), "http.request")
	defer span.End()

	span.SetAttributes(
		attribute.String("http.method", r.Method),
		attribute.String("http.route", r.URL.Path),
	)

	m.next.ServeHTTP(w, r.WithContext(ctx))
}

func WithTracing(next http.Handler) http.Handler {
	return tracingMiddleware{next: next}
}
