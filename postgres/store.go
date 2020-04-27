package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// NewStore ...
func NewStore(dataSourceName string) (*Store, error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return &Store{
		UserStore: &UserStore{DB: db},
	}, nil
}

// Store ...
type Store struct {
	*UserStore
}
