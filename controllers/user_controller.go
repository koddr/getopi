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
		c.Status(500).JSON(fiber.Map{"error": true, "description": err.Error()})
		return
	}

	id, err := uuid.Parse(c.Params("uuid"))
	if err != nil {
		c.Status(500).JSON(fiber.Map{"error": true, "description": err.Error()})
		return
	}

	user, err := store.User(id)
	if err != nil {
		c.Status(500).JSON(fiber.Map{"error": true, "description": err.Error()})
		return
	}

	c.JSON(fiber.Map{"error": false, "description": "ok", "user": user})
}

// UsersController ...
func UsersController(c *fiber.Ctx) {
	store, err := postgres.OpenStore()
	if err != nil {
		c.Status(500).JSON(fiber.Map{"error": true, "description": err.Error()})
		return
	}

	users, err := store.Users()
	if err != nil {
		c.Status(500).JSON(fiber.Map{"error": true, "description": err.Error()})
		return
	}

	c.JSON(fiber.Map{"error": false, "description": "ok", "users": users})
}

// UserCreateController ...
func UserCreateController(c *fiber.Ctx) {
	store, err := postgres.OpenStore()
	if err != nil {
		c.Status(500).JSON(fiber.Map{"error": true, "description": err.Error()})
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
		c.Status(500).JSON(fiber.Map{"error": true, "description": err.Error()})
		return
	}

	c.JSON(fiber.Map{"error": false, "description": "ok"})
}
