package httpserver

import (
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// Server wraps the standard http.Server with logging and graceful shutdown support
type Server struct {
	srv    *http.Server
	logger *zap.Logger
}

// Config holds configuration for the HTTP server
type Config struct {
	Host         string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// New creates a new HTTP server with the given handler and config
func New(handler http.Handler, cfg Config, logger *zap.Logger) *Server {
	if cfg.ReadTimeout == 0 {
		cfg.ReadTimeout = 15 * time.Second
	}
	if cfg.WriteTimeout == 0 {
		cfg.WriteTimeout = 15 * time.Second
	}
	if cfg.IdleTimeout == 0 {
		cfg.IdleTimeout = 60 * time.Second
	}

	addr := cfg.Host + ":" + cfg.Port
	srv := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	return &Server{
		srv:    srv,
		logger: logger,
	}
}

// Start starts the HTTP server in a non-blocking manner
func (s *Server) Start() error {
	s.logger.Info("starting HTTP server", zap.String("addr", s.srv.Addr))
	return s.srv.ListenAndServe()
}

// Stop gracefully shuts down the server with a timeout
func (s *Server) Stop(ctx context.Context) error {
	s.logger.Info("stopping HTTP server")
	return s.srv.Shutdown(ctx)
}

// Addr returns the server's listening address
func (s *Server) Addr() string {
	return s.srv.Addr
}
