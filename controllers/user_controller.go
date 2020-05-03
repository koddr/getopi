package controllers

import (
	"time"

	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	"github.com/koddr/getopi/models"
	"github.com/koddr/getopi/postgres"
)

// UserController ...
func UserController(c *fiber.Ctx) {
	store, err := postgres.OpenStore()
	if err != nil {
		// DB connection error
		c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
		return
	}

	id, err := uuid.Parse(c.Params("uuid"))
	if err != nil {
		// Wrong UUID format
		c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
		return
	}

	user, err := store.User(id)
	if err != nil {
		// Not found
		c.Status(404).JSON(fiber.Map{"error": false, "msg": err.Error()})
		return
	}

	// OK result
	c.JSON(fiber.Map{"error": false, "msg": "ok", "user": user})
}

// UsersController ...
func UsersController(c *fiber.Ctx) {
	store, err := postgres.OpenStore()
	if err != nil {
		// DB connection error
		c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
		return
	}

	users, err := store.Users()
	if err != nil {
		// Not found
		c.Status(404).JSON(fiber.Map{"error": false, "msg": err.Error()})
		return
	}

	// OK result
	c.JSON(fiber.Map{"error": false, "msg": "ok", "users": users})
}

// UserCreateController ...
func UserCreateController(c *fiber.Ctx) {
	store, err := postgres.OpenStore()
	if err != nil {
		// DB connection error
		c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
		return
	}

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
		// Not inserted new user to DB
		c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
		return
	}

	// OK result
	c.JSON(fiber.Map{"error": false, "msg": "ok"})
}
