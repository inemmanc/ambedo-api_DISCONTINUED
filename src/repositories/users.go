package repositories

import (
	"ambedo-api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

// NewUserRepo creates a new user repository
func NewUserRepo(db *sql.DB) *users {
	return &users{db}
}

// CreateUser inserts a new user in the database
func (repo users) CreateUser(user models.DefaultUser) (uint64, error) {
	return 0, nil
}

