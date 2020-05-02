package models

import (
	"time"

	"github.com/google/uuid"
)

// Project ...
type Project struct {
	ID            uuid.UUID    `json:"id"`
	AuthorID      uuid.UUID    `json:"author_id"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	Alias         string       `json:"alias"`
	ProjectStatus int          `json:"project_status"`
	ProjectAttrs  ProjectAttrs `json:"project_attrs"`
}

// ProjectAttrs ...
type ProjectAttrs struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Reward      int      `json:"reward"`
	Quantity    int      `json:"quantity"`
	IsPrivate   bool     `json:"is_private"`
	Links       []string `json:"links"`
	Goals       []string `json:"goals"`
	Tags        []string `json:"tags"`
	Skills      []string `json:"skills"`
}

// ProjectStore ...
type ProjectStore interface {
	Project(id uuid.UUID) (Project, error)
	Projects() ([]Project, error)
	CreateProject(p *Project) error
	UpdateProject(p *Project) error
	DeleteProject(id uuid.UUID) error
}
