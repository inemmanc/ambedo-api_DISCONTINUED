package models

import (
	"ambedo-api/src/constants"
	"errors"
	"strings"
	"time"
)

// DefaultUser struct represents the default API user model
type DefaultUser struct {
	ID         uint64    `json:"id,omitempty"`
	Username   string    `json:"username,omitempty"`
	Name       string    `json:"name,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	JoinedDate time.Time `json:"joinedate,omitempty"`
}

// Prepare calls methods to validate and format received user
func (user *DefaultUser) Prepare(method string) error {
	if err := user.validate(method); err != nil {
		return err
	}
	user.format()
	return nil
}

func (user *DefaultUser) validate(method string) error {
	if user.Username == "" {
		return errors.New("o campo Username é obrigatório")
	}
	if user.Name == "" {
		return errors.New("o campo Name é obrigatório")
	}
	if user.Email == "" {
		return errors.New("o campo Email é obrigatório")
	}
	if method == constants.MethodRegister && user.Password == "" {
		return errors.New("o campo Password é obrigatório")
	}
	return nil
}

func (user *DefaultUser) format() {
	user.Username = strings.Replace(user.Username, " ", "", -1)
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
}
