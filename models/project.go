package models

import (
	"time"

	"github.com/google/uuid"
)

// Project ...
type Project struct {
	// Main info
	ID          uuid.UUID `db:"id" json:"id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	Status      string    `db:"status" json:"status"`
	IsPrivate   bool      `db:"is_private" json:"is_private"`
	Alias       string    `db:"alias" json:"alias"`
	Title       string    `db:"title" json:"title"`
	Reward      int       `db:"reward" json:"reward"`
	Quantity    int       `db:"quantity" json:"quantity"`
	Description string    `db:"description" json:"description"`
	Links       []string  `db:"links" json:"links"`
	Goals       []string  `db:"goals" json:"goals"`
	Tags        []string  `db:"tags" json:"tags"`
	Skills      []string  `db:"skills" json:"skills"`

	// Author
	AuthorID uuid.UUID `db:"author_id" json:"author_id"`
}

// ProjectStore ...
type ProjectStore interface {
	Project(id uuid.UUID) (Project, error)
	Projects() ([]Project, error)
	CreateProject(p *Project) error
	UpdateProject(p *Project) error
	DeleteProject(id uuid.UUID) error
}
