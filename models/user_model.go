package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// User ...
type User struct {
	ID           uuid.UUID `db:"id" json:"id" validate:"required,id"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	Email        string    `db:"email" json:"email" validate:"required,email"`
	PasswordHash string    `db:"password_hash" json:"password_hash,omitempty"`
	Username     string    `db:"username" json:"username" validate:"required,username"`
	UserStatus   int       `db:"user_status" json:"user_status"`
	UserAttrs    UserAttrs `db:"user_attrs" json:"user_attrs"`
}

// UserAttrs ...
type UserAttrs struct {
	IsPrivate bool              `json:"is_private"`
	Picture   string            `json:"picture"`
	FirstName string            `json:"first_name"`
	LastName  string            `json:"last_name"`
	About     string            `json:"about"`
	Links     map[string]string `json:"links"`
	Skills    []string          `json:"skills"`
}

// UserMethods ...
type UserMethods interface {
	FindUserByID(id uuid.UUID) (User, error)
	FindUserByUsername(username string) (User, error)
	GetUsers() ([]User, error)
	CreateUser(u *User) error
	UpdateUser(u *User) error
	DeleteUser(id uuid.UUID) error
}

// Value ...
// Make the UserAttrs struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (a UserAttrs) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan ...
// Make the UserAttrs struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (a *UserAttrs) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
