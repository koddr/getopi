package postgres

import (
	"github.com/jmoiron/sqlx"
)

// ProjectStore ...
type ProjectStore struct {
	*sqlx.DB
}
