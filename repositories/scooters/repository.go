package scooters

import (
	"context"
	"database/sql"

	"github.com/fguy/scooters-api/entities"
	"go.uber.org/fx"
)

const (
	queryAvailable = `
	SELECT 
		id, lat, lng 
	FROM 
		scooters 
	WHERE 
		earth_box(ll_to_earth($1, $2), $3) @> ll_to_earth(lat, lng) 
		AND earth_distance(ll_to_earth($1, $2), ll_to_earth(lat, lng)) < $3 
		AND is_reserved = FALSE`
	queryReserve = `UPDATE scooters SET is_reserved = TRUE WHERE id = $1`
)

type repository struct {
	db *sql.DB
}

func (r *repository) GetAvailable(ctx context.Context, lat, lng, radius float32) ([]*entities.Scooter, error) {
	rows, err := r.db.QueryContext(ctx, queryAvailable, lat, lng, radius)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*entities.Scooter

	for rows.Next() {
		scooter := &entities.Scooter{}
		if err := rows.Scan(
			&scooter.ID,
			&scooter.Lat,
			&scooter.Lng,
		); err != nil {
			return nil, err
		}
		result = append(result, scooter)
	}

	return result, nil
}

func (r *repository) Reserve(ctx context.Context, ID int32) error {
	_, err := r.db.ExecContext(ctx, queryReserve, ID)
	return err
}

// New -
func New(lc fx.Lifecycle, openDB func() (*sql.DB, error)) (Interface, error) {
	db, err := openDB()
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			return db.Close()
		},
	})

	return &repository{
		db: db,
	}, nil
}
