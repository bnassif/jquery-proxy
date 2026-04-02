package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/bnassif/jquery-proxy/pkg/client"
	"github.com/bnassif/jquery-proxy/pkg/config"

	"github.com/bnassif/jquery-proxy/pkg/server/params"
	"github.com/bnassif/jquery-proxy/pkg/server/query"
)

type Server struct {
	server *http.Server
	client *client.Client
	config *config.ServerConfig
	Logger *slog.Logger
}

func (s *Server) handleRequest(w http.ResponseWriter, r *http.Request) {
	s.Logger.Debug(
		"request received",
		slog.String("remote_address", r.RemoteAddr),
		slog.String("url", r.URL.String()),
	)

	// Build query parameters based on the request
	p, err := params.NewParams(r.URL)
	if err != nil {
		s.Logger.Error(
			"http request error",
			slog.String("url", r.URL.String()),
			slog.Int("status_code", http.StatusBadRequest),
			slog.Any("error", err),
		)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.Logger.Debug(
		"parsed query parameters",
		slog.String("url", r.URL.String()),
		slog.Any("parameters", p),
	)

	// Make the request
	resBody, err := s.client.Request(p.URL.String())
	if err != nil {
		internalError(r.URL, w, err, s.Logger)
		return
	}

	// Query the response, if specified and return a string
	q := query.NewQuery(resBody, p)
	resp, err := q.Run()
	if err != nil {
		internalError(r.URL, w, err, s.Logger)
		return
	}

	// Write the response
	_, err = w.Write([]byte(resp))
	if err != nil {
		internalError(r.URL, w, err, s.Logger)
		return
	}
	return
}

func (s *Server) Run() error {
	fmt.Printf("Running server on %s\n", s.server.Addr)
	s.Logger.Info(
		"server starting",
		slog.String("listen", s.server.Addr),
		slog.String("address", s.config.Address),
		slog.String("port", s.config.Port),
	)

	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	fmt.Println("Stopping server")
	s.Logger.Info("server shutting down")

	return s.server.Shutdown(ctx)
}
