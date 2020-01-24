package scooters

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	repo "github.com/fguy/scooters-api/repositories/scooters"
	"go.uber.org/zap"
)

// AvailableHandler -
type AvailableHandler struct {
	logger     *zap.Logger
	repository repo.Interface
}

// ServeHTTP implements http.Handler
func (h *AvailableHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	h.logger.Info("request", zap.Any("query", q))

	lat, err := strconv.ParseFloat(q.Get("lat"), 64)
	if err != nil {
		h.logger.Error("invalid lat", zap.Error(err))
		w.WriteHeader(400)
		return
	}
	lng, err := strconv.ParseFloat(q.Get("lng"), 64)
	if err != nil {
		h.logger.Error("invalid lng", zap.Error(err))
		w.WriteHeader(400)
		return
	}
	radius, err := strconv.ParseFloat(q.Get("radius"), 64)
	if err != nil {
		h.logger.Error("invalid radius", zap.Error(err))
		w.WriteHeader(400)
		return
	}

	scooters, err := h.repository.GetAvailable(context.Background(), float32(lat), float32(lng), float32(radius))
	if err != nil {
		h.logger.Error("error", zap.Error(err))
		w.WriteHeader(500)
	}
	h.logger.Debug("response", zap.Any("scooters", scooters))
	json.NewEncoder(w).Encode(scooters)
}

// NewAvailableHandler -
func NewAvailableHandler(logger *zap.Logger, repository repo.Interface) *AvailableHandler {
	logger.Info("executing NewAvailableHandler")
	return &AvailableHandler{
		logger:     logger,
		repository: repository,
	}
}
