package models

import (
	"time"
)

// Project ...
type Project struct {
	// Main info
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Status      string    `json:"status"`
	IsPrivate   bool      `json:"is_private"`
	Name        string    `json:"name"`
	Reward      int       `json:"reward"`
	Description string    `json:"description"`
	Links       []string  `json:"links"`
	Goals       []string  `json:"goals"`
	Tags        []string  `json:"tags"`

	// Author
	AuthorID int `json:"author_id"`

	// Needed skills
	Skills []int `json:"skills"`

	// Tasks
	Tasks []int `json:"tasks"`
}
