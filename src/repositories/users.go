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
	statement, err := repo.db.Prepare("INSERT INTO users (username, name, email, password) VALUES(?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(user.Username, user.Name, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}
