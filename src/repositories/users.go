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

func (u users) CreateUser(user models.DefaultUser) (uint64, error) {
	return 0, nil
}

