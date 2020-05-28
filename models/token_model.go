package models

import (
	"time"

	"github.com/google/uuid"
)

// Token ...
type Token struct {
	ID          uuid.UUID `db:"id" json:"refresh_token"`
	UserID      uuid.UUID `db:"user_id" json:"-"`
	CreatedAt   time.Time `db:"created_at" json:"-"`
	ExpiredAt   time.Time `db:"expired_at" json:"expired_at"`
	AccessToken string    `db:"access_token" json:"access_token"`
}

// TokenStore ...
type TokenStore interface {
	CreateToken(t *Token) (Token, error)
	FindTokenByID(id uuid.UUID) (Token, error)
	RefreshTokenByID(id uuid.UUID, t *Token) error
}
