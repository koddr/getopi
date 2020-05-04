package stores

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/koddr/getopi/utils"
	_ "github.com/lib/pq" // ...
)

// Store ...
//
// TODO: Add description
//
type Store struct {
	*UserStore
	*ProjectStore
}

// OpenStore ...
//
// TODO: Add description
//
func OpenStore() (*Store, error) {
	db, err := sqlx.Open("postgres", utils.GetDotEnvValue("POSTGRES_SERVER_URL"))
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
