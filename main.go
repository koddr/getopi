package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/google/uuid"
	"github.com/koddr/getopi/models"
	"github.com/koddr/getopi/postgres"
)

func main() {
	// Fiber app
	app := fiber.New()

	// DB Store
	store, err := postgres.NewStore("host=localhost dbname=koddr sslmode=disable")
	if err != nil {
		log.Fatal("error opening database")
	}

	// Settings
	loggerConfig := logger.Config{
		Format:     "${time} - ${method} ${path}\n",
		TimeFormat: "Mon, 2 Jan 2006 15:04:05 MST",
	}

	// Logger
	app.Use(logger.New(loggerConfig))

	// Routes
	app.Post("/user", func(c *fiber.Ctx) {
		// id := uuid.MustParse("62aa9a05-a329-43c5-b864-3f36b27e5888")

		if err := store.CreateUser(
			&models.User{
				ID:           uuid.New(),
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Time{},
				Email:        "example@example.com",
				PasswordHash: "secret",
				Username:     "example",
				UserStatus:   1,
				UserAttrs: models.UserAttrs{
					FirstName: "John",
				},
			},
		); err != nil {
			c.JSON(fiber.Map{
				"error":       true,
				"description": err.Error(),
			})
			c.Status(500)
			return
		}

		c.JSON(fiber.Map{
			"error":       false,
			"description": "ok",
		})
	})

	app.Get("/users", func(c *fiber.Ctx) {
		users, err := store.Users()
		if err != nil {
			c.Status(500)
		}

		c.JSON(fiber.Map{"users": users})
	})

	app.Listen(":3000")
}
