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
	Quantity    int       `json:"quantity"`
	Description string    `json:"description"`
	Links       []string  `json:"links"`
	Goals       []string  `json:"goals"`
	Tags        []string  `json:"tags"`

	// Author
	AuthorID int `json:"author_id"`

	// Skills list
	Skills []int `json:"skills"`

	// Tasks list
	Tasks []int `json:"tasks"`

	// Opinions list
	Opinions []int `json:"opinions"`
}
