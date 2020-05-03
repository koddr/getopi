package stores

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // ...
)

// OpenStore ...
func OpenStore() (*Store, error) {
	db, err := sqlx.Open("postgres", "host=localhost dbname=koddr sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return &Store{
		UserStore:    &UserStore{DB: db},
		ProjectStore: &ProjectStore{DB: db},
	}, nil
}

// Store ...
type Store struct {
	*UserStore
	*ProjectStore
}
