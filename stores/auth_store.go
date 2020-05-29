package stores

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/koddr/getopi/models"
)

// AuthStore ...
//
// TODO: Add description
//
type AuthStore struct {
	*sqlx.DB
}

// CreateResetPasswordIssue ...
// Insert new restore password issue
func (s *AuthStore) CreateResetPasswordIssue(r *models.ResetCode) error {
	if _, err := s.Exec(
		`INSERT INTO reset_codes VALUES ($1, $2, $3)`,
		r.ID,
		r.UserID,
		r.ResetCode,
	); err != nil {
		return err
	}
	return nil
}

// FindResetPasswordIssueByCode ...
//
// TODO: Add description
//
func (s *AuthStore) FindResetPasswordIssueByCode(code string) error {
	var resetCode models.ResetCode
	if err := s.Get(&resetCode, `SELECT * FROM reset_codes WHERE reset_code = $1`, code); err != nil {
		return err
	}
	return nil
}

// DeleteResetPasswordIssueByCode ...
// Delete restore password issue by code
func (s *AuthStore) DeleteResetPasswordIssueByCode(code string) error {
	if _, err := s.Exec(`DELETE FROM reset_codes WHERE reset_code = $1`, code); err != nil {
		return err
	}
	return nil
}

// DeleteResetPasswordIssueByUserID ...
// Delete restore password issue by User ID
func (s *AuthStore) DeleteResetPasswordIssueByUserID(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM reset_codes WHERE user_id = $1`, id); err != nil {
		return err
	}
	return nil
}
