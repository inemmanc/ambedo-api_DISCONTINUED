package repositories

import (
	"ambedo-api/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

// NewUserRepo creates a new user repository
func NewUserRepo(db *sql.DB) *users {
	return &users{db}
}

func (repo users) FindUsers(nameOrUsername string) ([]models.DefaultUser, error) {
	nameOrUsername = fmt.Sprintf("%%%s%%", nameOrUsername)

	rows, err := repo.db.Query(
		"SELECT id, username, name, email, joineddate FROM users WHERE username LIKE ? OR name LIKE ?",
		nameOrUsername, nameOrUsername,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.DefaultUser
	for rows.Next() {
		var user models.DefaultUser

		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Name,
			&user.Email,
			&user.JoinedDate,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
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
