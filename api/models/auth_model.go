package models

import (
	"github.com/google/uuid"
)

// Auth ...
type Auth struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// ForgetPassword ...
type ForgetPassword struct {
	Email string `json:"email" validate:"required,email"`
}

// ResetCode ...
type ResetCode struct {
	ID        uuid.UUID `db:"id" json:"-"`
	UserID    uuid.UUID `db:"user_id" json:"-"`
	ResetCode string    `db:"reset_code" json:"reset_code" validate:"required,reset_code"`
}

// AuthStore ...
type AuthStore interface {
	CreateResetPasswordIssue(r *ResetCode) error
	FindResetPasswordIssueByCode(c string) error
	DeleteResetPasswordIssueByCode(c string) error
}
