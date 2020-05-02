package models

import (
	"time"

	"github.com/google/uuid"
)

// Opinion ...
type Opinion struct {
	// Main info
	ID        uuid.UUID `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	Text      string    `db:"text" json:"text"`
	Pictures  []string  `db:"pictures" json:"pictures"`

	// Pointer to task
	TaskID uuid.UUID `db:"task_id" json:"task_id"`
}

// OpinionStore ...
type OpinionStore interface {
	Opinion(id uuid.UUID) (Opinion, error)
	OpinionsByTask(taskID uuid.UUID) ([]Opinion, error)
	CreateOpinion(o *Opinion) error
	UpdateOpinion(o *Opinion) error
	DeleteOpinion(id uuid.UUID) error
}
