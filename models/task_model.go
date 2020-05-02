package models

import (
	"time"

	"github.com/google/uuid"
)

// Task ...
type Task struct {
	ID         uuid.UUID `son:"id"`
	ProjectID  uuid.UUID `json:"project_id"`
	CreatedAt  time.Time `json:"created_at"`
	TaskStatus int       `json:"task_status"`
	TaskAttrs  TaskAttrs `json:"task_attrs"`
}

// TaskAttrs ...
type TaskAttrs struct {
	Priority    int    `json:"priority"`
	Model       int    `json:"model"`
	IsSkip      bool   `json:"is_skip"`
	Description string `json:"description"`
}

// TaskStore ...
type TaskStore interface {
	Task(id uuid.UUID) (Task, error)
	TasksByProject(projectID uuid.UUID) ([]Task, error)
	CreateTask(t *Task) error
	UpdateTask(t *Task) error
	DeleteTask(id uuid.UUID) error
}
