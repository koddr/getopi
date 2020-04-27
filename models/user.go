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
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Attrs        Attrs     `json:"attrs"`
}

// UserMethods ...
type UserMethods interface {
	User(id uuid.UUID) (User, error)
	Users() ([]User, error)
	CreateUser(u *User) error
	UpdateUser(u *User) error
	DeleteUser(id uuid.UUID) error
}

// Attrs ...
type Attrs struct {
	Status    string            `json:"status"`
	IsPrivate bool              `json:"is_private"`
	Username  string            `json:"username"`
	Picture   string            `json:"picture"`
	FirstName string            `json:"first_name"`
	LastName  string            `json:"last_name"`
	About     string            `json:"about"`
	Links     map[string]string `json:"links"`
	Skills    []string          `json:"skills"`
}

// Value ...
// Make the Attrs struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (a Attrs) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan ...
// Make the Attrs struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (a *Attrs) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
