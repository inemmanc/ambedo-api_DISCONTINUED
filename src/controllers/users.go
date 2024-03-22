package controllers

import (
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
	w.Write([]byte("creating user func"))
}

// UpdateUser updates a specific user information in the database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user func"))
}

// DeleteUser deletes a specific user information in the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user func"))
}
