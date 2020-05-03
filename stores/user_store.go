package stores

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/koddr/getopi/models"
)

// UserStore ...
type UserStore struct {
	*sqlx.DB
}

// User ...
func (s *UserStore) User(username string) (models.User, error) {
	var user models.User
	if err := s.Get(&user, `SELECT * FROM users WHERE username = $1`, username); err != nil {
		return models.User{}, err
	}
	return user, nil
}

// Users ...
func (s *UserStore) Users() ([]models.User, error) {
	var users []models.User
	if err := s.Select(&users, `SELECT * FROM users`); err != nil {
		return []models.User{}, err
	}
	return users, nil
}

// CreateUser ...
func (s *UserStore) CreateUser(u *models.User) error {
	if _, err := s.Exec(
		`INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		u.ID,
		u.CreatedAt,
		u.UpdatedAt,
		u.Email,
		u.PasswordHash,
		u.Username,
		u.UserStatus,
		u.UserAttrs,
	); err != nil {
		return err
	}
	return nil
}

// UpdateUser ...
func (s *UserStore) UpdateUser(u *models.User) error {
	// Update user fields
	if _, err := s.Exec(
		`UPDATE users SET updated_at = $1, email = $2, password_hash = $3, username = $4, user_attrs = $5 WHERE id = $6`,
		u.UpdatedAt,
		u.Email,
		u.PasswordHash,
		u.Username,
		u.UserAttrs,
		u.ID,
	); err != nil {
		return err
	}

	// OK Result
	return nil
}

// DeleteUser ...
func (s *UserStore) DeleteUser(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM users WHERE id = $1`, id); err != nil {
		return err
	}
	return nil
}
