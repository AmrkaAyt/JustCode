package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"lecture10/config"
)

func NewDBConnection(configPath string) (*sql.DB, error) {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return nil, err
	}

	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
	)

	db, err := sql.Open(cfg.Database.Driver, connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
