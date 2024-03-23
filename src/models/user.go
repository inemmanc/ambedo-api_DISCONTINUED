package models

import "time"

// DefaultUser struct represents the default API user model
type DefaultUser struct {
	ID         uint64    `json:"id,omitempty"`
	Username   string    `json:"username,omitempty"`
	Name       string    `json:"name,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	JoinedDate time.Time `json:"joinedate,omitempty"`
}
