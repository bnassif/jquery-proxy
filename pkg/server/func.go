package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/bnassif/jquery-proxy/pkg/client"
	"github.com/bnassif/jquery-proxy/pkg/config"
	"github.com/bnassif/jquery-proxy/pkg/logging"
)

func NewServer(cfg *config.Config) *Server {
	baseLogger := logging.New(cfg.Logging, "jquery-proxy")

	s := &Server{
		config: &cfg.Server,
		client: client.NewClient(&cfg.Client, baseLogger),
		Logger: baseLogger.With(slog.String("component", "server")),
	}

	addrPortStr := fmt.Sprintf("%s:%s", cfg.Server.Address, cfg.Server.Port)

	s.server = &http.Server{
		Addr:         addrPortStr,
		ReadTimeout:  cfg.Server.GetReadTimeout(),
		WriteTimeout: cfg.Server.GetWriteTimeout(),
		IdleTimeout:  cfg.Server.GetIdleTimeout(),
		Handler:      http.HandlerFunc(s.handleRequest),
	}

	return s
}
