package models

import (
	"time"

	"github.com/google/uuid"
)

// Task ...
type Task struct {
	// Main info
	ID          uuid.UUID `db:"id" json:"id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	Skip        bool      `db:"skip" json:"skip"`
	Status      string    `db:"status" json:"status"`
	Model       string    `db:"model" json:"model"`
	Description string    `db:"description" json:"description"`

	// Pointer to project
	ProjectID uuid.UUID `db:"project_id" json:"project_id"`
}

// TaskStore ...
type TaskStore interface {
	Task(id uuid.UUID) (Task, error)
	TasksByProject(projectID uuid.UUID) ([]Task, error)
	CreateTask(t *Task) error
	UpdateTask(t *Task) error
	DeleteTask(id uuid.UUID) error
}
