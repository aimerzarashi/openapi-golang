package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func New() (*sql.DB, error) {
	driverName := os.Getenv("DB_DRIVER")
	if driverName == "" {
		driverName = "postgres"
	}
	dataSourceName := os.Getenv("DB_DSN")
	if dataSourceName == "" {
		dataSourceName = "host=localhost port=5432 user=user password=password dbname=openapi sslmode=disable"
	}

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
			return nil, err
	}

	return db, nil
}