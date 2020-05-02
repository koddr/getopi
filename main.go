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

	// Connect to DB
	s, err := postgres.NewStore("host=localhost dbname=koddr sslmode=disable")
	if err != nil {
		log.Fatal("error opening database")
	}

	// Settings
	loggerConfig := logger.Config{
		Format:     "${time} [${status}] ${method} ${path}\n",
		TimeFormat: "Mon, 2 Jan 2006 15:04:05 MST",
	}

	// Logger
	app.Use(logger.New(loggerConfig))

	// POST
	app.Post("/user", func(c *fiber.Ctx) {
		if err := s.CreateUser(
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
			c.Status(500).JSON(fiber.Map{"error": true, "description": err.Error()})
			return
		}

		c.JSON(fiber.Map{"error": false, "description": "ok"})
	})

	// GET by UUID
	app.Get("/user/:uuid", func(c *fiber.Ctx) {
		id := uuid.MustParse(c.Params("uuid"))
		user, err := s.User(id)
		if err != nil {
			c.Status(500).JSON(fiber.Map{"error": true, "description": err.Error()})
			return
		}

		c.JSON(fiber.Map{"error": false, "description": "ok", "user": user})
	})

	// GET All
	app.Get("/users", func(c *fiber.Ctx) {
		users, err := s.Users()
		if err != nil {
			c.Status(500).JSON(fiber.Map{"error": true, "description": err.Error()})
			return
		}

		c.JSON(fiber.Map{"error": false, "description": "ok", "users": users})
	})

	app.Listen(":3000")
}
