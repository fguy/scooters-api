package scooters

import (
	"context"

	"github.com/fguy/scooters-api/entities"
)

//go:generate mockgen -destination=../../mocks/repositories/scooters/interface.go github.com/fguy/scooters-api/repositories/scooters Interface

// Interface is a interface of scooters repository
type Interface interface {
	GetAvailable(context.Context, float32, float32, float32) ([]*entities.Scooter, error)
	Reserve(context.Context, int32) error
}
