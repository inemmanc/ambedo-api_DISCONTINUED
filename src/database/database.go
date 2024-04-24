package database

import (
	"ambedo-api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// opens the DB connection and returns it
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DefaultConnectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
