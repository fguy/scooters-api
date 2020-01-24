package factories

import (
	"github.com/fguy/scooters-api/entities"
)

var (
	// Scooters -
	Scooters = []*entities.Scooter{
		{
			ID:  10,
			Lat: 37.788548,
			Lng: -122.411548,
		},
		{
			ID:  8,
			Lat: 37.783223,
			Lng: -122.398630,
		},
	}
)
