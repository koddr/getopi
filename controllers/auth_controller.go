package controllers

import (
	"github.com/gofiber/fiber"
	"github.com/koddr/getopi/stores"
	"github.com/koddr/getopi/utils"
)

// Authentication ...
func Authentication(c *fiber.Ctx) {
	// Struct for login and password
	type authData struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	// Create new User struct
	auth := &authData{}

	// Create new validator
	validate := utils.Validate("auth")

	// Check received JSON data
	if errBodyParser := c.BodyParser(auth); errBodyParser != nil {
		// Incorrect data
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errBodyParser.Error()})
		return
	}

	// Check fields validation
	if errValidate := validate.Struct(auth); errValidate != nil {
		// Return invalid fields
		c.Status(500).JSON(fiber.Map{"error": true, "msg": utils.ValidateErrors(errValidate)})
		return
	}

	// Create DB connection
	db, errConnectDB := stores.OpenStore()
	if errConnectDB != nil {
		// DB connection error
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errConnectDB.Error()})
		return
	}

	// Find user by email
	user, errFindUserByEmail := db.FindUserByEmail(auth.Email)
	if errFindUserByEmail != nil {
		// User not found
		c.Status(404).JSON(fiber.Map{"error": true, "msg": "user not found", "user": nil})
		return
	}

	if utils.ComparePasswords(user.PasswordHash, auth.Password) {
		// Create JWT token
		token, errGenerateJWT := utils.GenerateJWT("user", user.ID.String())
		if errGenerateJWT != nil {
			// Fail create JWT token
			c.Status(500).JSON(fiber.Map{"error": true, "msg": errGenerateJWT.Error()})
			return
		}

		c.JSON(fiber.Map{"error": true, "msg": nil, "token": token})
	} else {
		// Fail authentication
		c.Status(401).JSON(fiber.Map{"error": true, "msg": "incorrect email or password"})
		return
	}
}
