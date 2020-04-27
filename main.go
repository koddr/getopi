package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/koddr/getopi/models"
)

func main() {
	db, err := pgx.Connect(context.Background(), "postgres://koddr@localhost/koddr?sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	// Initialize a new User struct and add some values.
	user := new(models.User)
	user.Email = "test@example.com"
	user.PasswordHash = "secret"

	// Initialize a new Attrs struct and add some values.
	attrs := new(models.Attrs)
	attrs.Username = "john_doe_2001"
	attrs.Skills = []string{"Basil", "Garlic", "Parmesan", "Pine nuts", "Olive oil"}

	// The database driver will call the Value() method and and marshall the
	// attrs struct to JSON before the INSERT.
	_, err = db.Exec(
		context.Background(),
		`INSERT INTO users (email, password_hash, attrs) VALUES ($1, $2, $3)`,
		user.Email,
		user.PasswordHash,
		attrs,
	)
	if err != nil {
		log.Fatalf("Can not insert user: %v\n", err)
	}

	// Similarly, we can also fetch data from the database, and the driver
	// will call the Scan() method to unmarshal the data to an Attr struct.
	err = db.QueryRow(
		context.Background(),
		`SELECT id, attrs FROM users ORDER BY id DESC LIMIT 1`,
	).Scan(&user.ID, &user.Attrs)
	if err != nil {
		log.Fatal(err)
	}

	// You can then use the struct fields as normal...
	log.Printf("ID: %d, Username: %s", user.ID, user.Attrs.Username)
}
