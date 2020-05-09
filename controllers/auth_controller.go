package controllers

import (
	"github.com/gofiber/fiber"
	"github.com/koddr/getopi/models"
	"github.com/koddr/getopi/utils"
)

// Auth ...
func Auth(c *fiber.Ctx) {
	// Create new User struct
	user := &models.User{}

	// Check received JSON data
	if err := c.BodyParser(user); err != nil {
		// Incorrect data
		c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
		return
	}

	// Check Email for empty value
	if user.Email == "" {
		// Incorrect data
		c.Status(401).JSON(fiber.Map{"error": true, "msg": "incorrect Email"})
		return
	}

	// Convert string to UUID
	token, _ := utils.GenerateJWT("user", user.ID)
	c.JSON(fiber.Map{"token": token})
}
