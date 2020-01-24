package scooters

import (
	"net/http"

	"github.com/fguy/scooters-api/config"

	"go.uber.org/zap"
)

// NewHTTPServer -
func NewHTTPServer(
	cfg *config.AppConfig,
	logger *zap.Logger,
	availableHandler *AvailableHandler,
	reserveHandler *ReserveHandler,
) *http.Server {
	http.Handle("/api/v1/scooters/available", availableHandler)
	http.Handle("/api/v1/scooters/reserve", reserveHandler)
	return &http.Server{Addr: cfg.Addr}
}
