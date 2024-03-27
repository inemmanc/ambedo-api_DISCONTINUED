package models

import (
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
func (user *DefaultUser) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}
	user.format()
	return nil
}

func (user *DefaultUser) validate() error {
	if user.Username == "" {
		return errors.New("o campo Username é obrigatório")
	}
	if user.Name == "" {
		return errors.New("o campo Name é obrigatório")
	}
	if user.Email == "" {
		return errors.New("o campo Email é obrigatório")
	}
	if user.Password == "" {
		return errors.New("o campo Password é obrigatório")
	}
	return nil
}

func (user *DefaultUser) format() {
	user.Username = strings.TrimSpace(user.Username)
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
}
