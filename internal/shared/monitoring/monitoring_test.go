package monitoring

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlerExportsMetrics(t *testing.T) {
	h := NewHandler()

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/metrics", nil)

	h.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", recorder.Code)
	}

	body := recorder.Body.String()
	if !strings.Contains(body, "http_requests_total") {
		t.Fatalf("expected metrics body to contain http_requests_total, got %q", body)
	}
}
