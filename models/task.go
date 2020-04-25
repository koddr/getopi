package models

import (
	"time"
)

// Task ...
type Task struct {
	// Main info
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Skip        bool      `json:"skip"`
	Status      string    `json:"status"`
	Model       string    `json:"model"`
	Description string    `json:"description"`

	// Pointer to project
	ProjectID int `json:"project_id"`
}
