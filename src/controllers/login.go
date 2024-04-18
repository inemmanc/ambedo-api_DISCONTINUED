package controllers

import (
	"ambedo-api/src/auth"
	"ambedo-api/src/database"
	"ambedo-api/src/models"
	"ambedo-api/src/repositories"
	"ambedo-api/src/responses"
	"ambedo-api/src/security"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Login is used for authenticate a user to the API
func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	var user models.DefaultUser
	if err := json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepo(db)
	dbSavedUser, err := repository.FindUserByEmail(user.Email)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err := security.VerifyPassword(dbSavedUser.Password, user.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	token, err := auth.CreateToken(dbSavedUser.ID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	// TEMP RESPONSE ---- REMOVE
	w.Write([]byte(token))
}
