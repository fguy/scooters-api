package repositories

import (
	"database/sql"

	"github.com/fguy/scooters-api/config"
	// load db driver
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

const driverName = "postgres"

// NewDB -
func NewDB(logger *zap.Logger, cfg *config.AppConfig) (func() (*sql.DB, error), error) {
	return func() (*sql.DB, error) {
		db, err := sql.Open(driverName, cfg.DSN)
		if err != nil {
			logger.Error("can not connnect to the database", zap.Error(err))
			return nil, err
		}
		return db, nil
	}, nil
}
