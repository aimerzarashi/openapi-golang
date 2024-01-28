package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func New() (*sql.DB, error) {
	dbDriver := "postgres"
	dsn := "host=openapi-db port=5432 user=user password=password dbname=openapi sslmode=disable"
	// dsn := "host=localhost port=5432 user=user password=password dbname=openapi sslmode=disable"

	db, openErr := sql.Open(dbDriver, dsn)
	if openErr != nil {
		return nil, openErr
	}

	return db, nil
}