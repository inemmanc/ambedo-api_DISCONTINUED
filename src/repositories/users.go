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

// FindUsers search for users based on name or username in the database
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

// FindUserByID returns a user in the database based on their user ID
func (repo users) FindUserByID(userID uint64) (models.DefaultUser, error) {
	rows, err := repo.db.Query(
		"SELECT id, username, name, email, joineddate FROM users WHERE id = ?",
		userID,
	)
	if err != nil {
		return models.DefaultUser{}, err
	}
	defer rows.Close()

	var user models.DefaultUser
	if rows.Next() {
		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Name,
			&user.Email,
			&user.JoinedDate,
		); err != nil {
			return models.DefaultUser{}, err
		}
	}
	return user, nil
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

// UpdateUser updates a user's information in the database
func (repo users) UpdateUser(userID uint64, user models.DefaultUser) error {
	statement, err := repo.db.Prepare(
		"UPDATE users SET username = ?, name = ?, email = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Username, user.Name, user.Email, userID); err != nil {
		return err
	}

	return nil
}

// DeleteUser deletes a user's information in the database
func (repo users) DeleteUser(userID uint64) error {
	statement, err := repo.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userID); err != nil {
		return err
	}

	return nil
}

// FindUserByEmail search for a user by their email and returns their ID and hased Password
func (repo users) FindUserByEmail(userEmail string) (models.DefaultUser, error) {
	row, err := repo.db.Query("SELECT id, password FROM users WHERE email = ?", userEmail)
	if err != nil {
		return models.DefaultUser{}, err
	}
	defer row.Close()

	var user models.DefaultUser
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Password); err != nil {
			return models.DefaultUser{}, err
		}
	}

	return user, nil
}
