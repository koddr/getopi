package models

import (
	"time"
)

// User ...
type User struct {
	// Main info
	ID        int               `json:"id"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at,omitempty"`
	Status    string            `json:"status"`
	IsPrivate bool              `json:"is_private"`
	Alias     string            `json:"alias"`
	Picture   string            `json:"picture"`
	FirstName string            `json:"first_name"`
	LastName  string            `json:"last_name"`
	About     string            `json:"about"`
	Links     map[string]string `json:"links"`

	// Login
	Email        string `json:"email"`
	PasswordHash string `json:"-"`

	// Skills
	Skills []int `json:"skills"`

	// Projects
	Projects []int `json:"projects"`
}
