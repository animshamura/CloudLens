package gateway

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTracingMiddlewareWrapsRequests(t *testing.T) {
	var called bool
	handler := tracingMiddleware{next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	})}

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	handler.ServeHTTP(recorder, req)

	if !called {
		t.Fatal("expected wrapped handler to run")
	}
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", recorder.Code)
	}
}
