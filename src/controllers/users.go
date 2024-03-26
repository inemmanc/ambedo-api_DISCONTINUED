package controllers

import (
	"ambedo-api/src/database"
	"ambedo-api/src/models"
	"ambedo-api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// FindUsers searchs for all users in the database
func FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("searching users func"))
}

// FindUser search for a specific user in the database
func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("searching a specif user func"))
}

// CreateUser creates a new user into the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		// TEMP RESPONSE (REMOVE) ----
		w.Write([]byte("Error"))
	}

	var user models.DefaultUser
	if err := json.Unmarshal(requestBody, &user); err != nil {
		// TEMP RESPONSE (REMOVE) ----
		w.Write([]byte("Error unmarshal"))
	}

	db, err := database.Connect()
	if err != nil {
		// TEMP RESPONSE (REMOVE) ----
		w.Write([]byte("Error database"))
	}

	repo := repositories.NewUserRepo(db)
	lastInsertID, err := repo.CreateUser(user)
	if err != nil {
		// TEMP RESPONSE (REMOVE) ----
		w.Write([]byte("Error repository"))
	}

	w.Write([]byte(fmt.Sprintf("last id inserted = %d", lastInsertID)))
}

// UpdateUser updates a specific user information in the database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user func"))
}

// DeleteUser deletes a specific user information in the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user func"))
}
