package stores

import (
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib" // ...
	"github.com/jmoiron/sqlx"
	"github.com/koddr/getopi/utils"
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
	db, err := sqlx.Connect("pgx", utils.GetDotEnvValue("POSTGRES_SERVER_URL"))
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
