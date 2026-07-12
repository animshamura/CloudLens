package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandlerReportsLivenessAndReadiness(t *testing.T) {
	h := NewHandler()

	ll := httptest.NewRecorder()
	lreq := httptest.NewRequest(http.MethodGet, "/health/live", nil)
	h.ServeHTTP(ll, lreq)
	require.Equal(t, http.StatusOK, ll.Code)

	rr := httptest.NewRecorder()
	rreq := httptest.NewRequest(http.MethodGet, "/health/ready", nil)
	h.ServeHTTP(rr, rreq)
	require.Equal(t, http.StatusOK, rr.Code)
}
