package models

import (
	"time"
)

// Opinion ...
type Opinion struct {
	// Main info
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Text      string    `json:"text"`
	Pictures  []string  `json:"pictures"`

	// Pointer to task
	TaskID int `json:"task_id"`

	// Pointer to user
	UserID int `json:"user_id"`
}
