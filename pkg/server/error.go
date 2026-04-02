package server

import (
	"log/slog"
	"net/http"
	"net/url"
)

func internalError(url *url.URL, writer http.ResponseWriter, err error, logger *slog.Logger) {
	logger.Error(
		"http internal error",
		slog.String("url", url.String()),
		slog.Int("status_code", http.StatusInternalServerError),
		slog.Any("error", err),
	)
	http.Error(writer, err.Error(), http.StatusInternalServerError)
}
