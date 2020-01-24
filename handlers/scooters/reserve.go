package scooters

import (
	"context"
	"net/http"
	"strconv"

	repo "github.com/fguy/scooters-api/repositories/scooters"
	"go.uber.org/zap"
)

// ReserveHandler -
type ReserveHandler struct {
	logger     *zap.Logger
	repository repo.Interface
}

// ServeHTTP implements http.Handler
func (h *ReserveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	h.logger.Info("request", zap.Any("query", q))

	ID, err := strconv.ParseInt(q.Get("id"), 10, 32)
	if err != nil {
		h.logger.Error("invalid lat", zap.Error(err))
		w.WriteHeader(400)
		return
	}

	err = h.repository.Reserve(context.Background(), int32(ID))
	if err != nil {
		h.logger.Error("error", zap.Error(err))
		w.WriteHeader(500)
	}
	w.WriteHeader(200)
}

// NewReserveHandler -
func NewReserveHandler(logger *zap.Logger, repository repo.Interface) *ReserveHandler {
	logger.Info("executing NewReserveHandler")
	return &ReserveHandler{
		logger:     logger,
		repository: repository,
	}
}
