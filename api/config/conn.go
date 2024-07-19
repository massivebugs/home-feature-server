package config

import (
	"database/sql"
	"fmt"
)

func CreateDatabaseConnection(cfg *Config) (*sql.DB, error) {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?multiStatements=true",
			cfg.DBUser,
			cfg.DBPassword,
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBDatabase,
		),
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
