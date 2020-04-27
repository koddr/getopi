package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/koddr/getopi/models"
)

// UserStore ...
type UserStore struct {
	*sqlx.DB
}

// User ...
func (s *UserStore) User(id uuid.UUID) (models.User, error) {
	var user models.User
	if err := s.Get(&user, `SELECT * FROM users WHERE id = $1`, id); err != nil {
		return models.User{}, fmt.Errorf("error getting user: %w", err)
	}
	return user, nil
}

// Users ...
func (s *UserStore) Users() ([]models.User, error) {
	var users []models.User
	if err := s.Select(&users, `SELECT * FROM users`); err != nil {
		return []models.User{}, fmt.Errorf("error getting user: %w", err)
	}
	return users, nil
}

// CreateUser ...
func (s *UserStore) CreateUser(u *models.User) error {
	if err := s.Get(
		u,
		`INSERT INTO users VALUES ($1, $2, $3, $4) RETURNING *`,
		u.ID,
		u.Email,
		u.PasswordHash,
		u.Attrs,
	); err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}
	return nil
}

// UpdateUser ...
func (s *UserStore) UpdateUser(u *models.User) error {
	if err := s.Get(
		u,
		`UPDATE users SET email = $1, attrs = $2 WHERE id = $3 RETURNING *`,
		u.Email,
		u.Attrs,
		u.ID,
	); err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}
	return nil
}

// DeleteUser ...
func (s *UserStore) DeleteUser(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM users WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}
	return nil
}
