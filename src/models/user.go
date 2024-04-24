package models

import (
	"ambedo-api/src/constants"
	"ambedo-api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// represents the default API user model
type DefaultUser struct {
	ID         uint64    `json:"id,omitempty"`
	Username   string    `json:"username,omitempty"`
	Name       string    `json:"name,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	JoinedDate time.Time `json:"joinedate,omitempty"`
}

// Prepare calls methods to validate and format received user information
func (user *DefaultUser) Prepare(method string) error {
	if err := user.validate(method); err != nil {
		return err
	}
	if err := user.format(method); err != nil {
		return err
	}
	return nil
}

func (user *DefaultUser) validate(method string) error {
	if user.Username == "" {
		return errors.New("you need to enter a valid Username")
	}
	if user.Name == "" {
		return errors.New("you need to enter a valid Name")
	}
	if user.Email == "" {
		return errors.New("you need to enter a valid Email")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("you inserted a invalid email")
	}

	if method == "register" && user.Password == "" {
		return errors.New("you need to enter a valid Password")
	}
	return nil
}

func (user *DefaultUser) format(method string) error {
	user.Username = strings.Replace(user.Username, " ", "", -1)
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)

	if method == constants.MethodRegisterUser {
		hashedPassword, err := security.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}

	return nil
}
