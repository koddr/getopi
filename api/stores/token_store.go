package stores

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/koddr/getopi/models"
)

// TokenStore ...
//
// TODO: Add description
//
type TokenStore struct {
	*sqlx.DB
}

// CreateToken ...
//
// TODO: Add description
//
func (s *TokenStore) CreateToken(t *models.Token) (models.Token, error) {
	// Insert new token
	if _, err := s.Exec(
		`INSERT INTO tokens VALUES ($1, $2, $3, $4, $5)`,
		t.ID,
		t.UserID,
		t.CreatedAt,
		t.ExpiredAt,
		t.AccessToken,
	); err != nil {
		return models.Token{}, err
	}

	// Return new token
	return models.Token{
		ID:          t.ID,
		UserID:      t.UserID,
		CreatedAt:   t.CreatedAt,
		ExpiredAt:   t.ExpiredAt,
		AccessToken: t.AccessToken,
	}, nil
}

// FindTokenByID ...
//
// TODO: Add description
//
func (s *TokenStore) FindTokenByID(id uuid.UUID) (models.Token, error) {
	var token models.Token
	if err := s.Get(&token, `SELECT * FROM tokens WHERE id = $1`, id); err != nil {
		return models.Token{}, err
	}
	return token, nil
}

// DeleteTokenByID ...
//
// TODO: Add description
//
func (s *TokenStore) DeleteTokenByID(id uuid.UUID) error {
	// Delete exists token
	if _, err := s.Exec(`DELETE FROM tokens WHERE id = $1`, id); err != nil {
		return err
	}
	return nil
}

// DeleteTokenByUserID ...
//
// TODO: Add description
//
func (s *TokenStore) DeleteTokenByUserID(id uuid.UUID) error {
	// Delete exists token
	if _, err := s.Exec(`DELETE FROM tokens WHERE user_id = $1`, id); err != nil {
		return err
	}
	return nil
}
