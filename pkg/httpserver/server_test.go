package httpserver

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestServerStart(t *testing.T) {
	logger := zap.NewNop()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test"))
	})

	cfg := Config{
		Host: "127.0.0.1",
		Port: "0", // Use any available port
	}

	srv := New(handler, cfg, logger)
	require.NotNil(t, srv)
	require.NotEmpty(t, srv.Addr())
}

func TestServerConfig(t *testing.T) {
	logger := zap.NewNop()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	cfg := Config{
		Host:         "127.0.0.1",
		Port:         "9999",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	srv := New(handler, cfg, logger)
	require.Equal(t, "127.0.0.1:9999", srv.Addr())
}

func TestServerDefaultTimeouts(t *testing.T) {
	logger := zap.NewNop()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	cfg := Config{
		Host: "127.0.0.1",
		Port: "9999",
	}

	srv := New(handler, cfg, logger)
	require.Equal(t, 15*time.Second, srv.srv.ReadTimeout)
	require.Equal(t, 15*time.Second, srv.srv.WriteTimeout)
	require.Equal(t, 60*time.Second, srv.srv.IdleTimeout)
}

func TestServerHandler(t *testing.T) {
	logger := zap.NewNop()
	expectedBody := "hello world"
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(expectedBody))
	})

	cfg := Config{
		Host: "127.0.0.1",
		Port: "9999",
	}

	srv := New(handler, cfg, logger)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)

	srv.srv.Handler.ServeHTTP(recorder, request)

	require.Equal(t, http.StatusOK, recorder.Code)
	require.Equal(t, expectedBody, recorder.Body.String())
}

func TestServerGracefulShutdown(t *testing.T) {
	logger := zap.NewNop()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	cfg := Config{
		Host: "127.0.0.1",
		Port: "0",
	}

	srv := New(handler, cfg, logger)

	// Test that Shutdown doesn't panic with a nil or stopped server
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := srv.Stop(ctx)
	// Shutdown on a server that hasn't been started returns http.ErrServerClosed
	require.True(t, err == nil || err == http.ErrServerClosed)
}
