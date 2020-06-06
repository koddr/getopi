package stores

import (
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib" // ...
	"github.com/jmoiron/sqlx"
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
	db, errConnect := sqlx.Connect("pgx", os.Getenv("POSTGRES_SERVER_URL"))
	if errConnect != nil {
		return nil, fmt.Errorf("error opening database: %w", errConnect)
	}

	// Connection settings
	maxConn, _ := strconv.Atoi(os.Getenv("POSTGRES_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("POSTGRES_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("POSTGRES_MAX_LIFETIME_CONNECTIONS"))

	db.SetMaxOpenConns(maxConn)                           // The default is 0 (unlimited)
	db.SetMaxIdleConns(maxIdleConn)                       // defaultMaxIdleConns = 2
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn)) // 0, connections are reused forever

	if errPing := db.Ping(); errPing != nil {
		db.Close()
		return nil, fmt.Errorf("error connecting to database: %w", errPing)
	}

	return &Store{
		AuthStore:    &AuthStore{DB: db},
		UserStore:    &UserStore{DB: db},
		ProjectStore: &ProjectStore{DB: db},
		TokenStore:   &TokenStore{DB: db},
	}, nil
}
