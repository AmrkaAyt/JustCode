package database

import (
	"Lecture9/config"
	"database/sql"
	_ "github.com/lib/pq"
)

func NewDBConnection() (*sql.DB, error) {
	cfg, err := config.LoadConfig("config/config.yml")
	if err != nil {
		return nil, err
	}

	connectionString := cfg.DatabaseURL
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
