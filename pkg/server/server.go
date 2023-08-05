package server

import (
	"context"
	"net/http"
	"time"
)

type Server interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

type server struct {
	srv *http.Server
}

func NewServer(addr string, handler http.Handler) Server {
	return &server{
		srv: &http.Server{
			Addr:         addr,
			Handler:      handler,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
}

func (s *server) ListenAndServe() error {
	return s.srv.ListenAndServe()
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
