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
	*AuthStore
	*UserStore
	*ProjectStore
	*TokenStore
}

// OpenStore ...
//
// TODO: Add description
//
func OpenStore() (*Store, error) {
	db, errConnect := sqlx.Connect("pgx", utils.GetDotEnvValue("POSTGRES_SERVER_URL"))
	if errConnect != nil {
		return nil, fmt.Errorf("error opening database: %w", errConnect)
	}

	if errPing := db.Ping(); errPing != nil {
		return nil, fmt.Errorf("error connecting to database: %w", errPing)
	}

	return &Store{
		AuthStore:    &AuthStore{DB: db},
		UserStore:    &UserStore{DB: db},
		ProjectStore: &ProjectStore{DB: db},
		TokenStore:   &TokenStore{DB: db},
	}, nil
}
