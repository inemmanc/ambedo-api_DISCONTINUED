package controllers

import (
	"ambedo-api/src/database"
	"ambedo-api/src/models"
	"ambedo-api/src/repositories"
	"ambedo-api/src/responses"
	"encoding/json"
	"fmt"
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

	repository := repositories.NewUserRepo(db)
	// TEMP -- MISSING repository func 
	fmt.Println(repository)
}
