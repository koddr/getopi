package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Project ...
type Project struct {
	ID            uuid.UUID    `db:"id" json:"id"`
	AuthorID      uuid.UUID    `db:"author_id" json:"author_id"`
	CreatedAt     time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time    `db:"updated_at" json:"updated_at"`
	Alias         string       `db:"alias" json:"alias"`
	ProjectStatus int          `db:"project_status" json:"project_status"`
	ProjectAttrs  ProjectAttrs `db:"project_attrs" json:"project_attrs"`
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

// Value make the ProjectAttrs struct implement the driver.Valuer interface.
// This method simply returns the JSON-encoded representation of the struct.
func (a ProjectAttrs) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan make the ProjectAttrs struct implement the sql.Scanner interface.
// This method simply decodes a JSON-encoded value into the struct fields.
func (a *ProjectAttrs) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
