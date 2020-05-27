package controllers

import (
	"time"

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

	// Check password
	if utils.ComparePasswords(user.PasswordHash, auth.Password) {
		// Create JWT access_token
		accessToken, errGenerateAccessJWT := utils.GenerateAccessJWT("user", user.ID.String())
		if errGenerateAccessJWT != nil {
			// Fail create JWT token
			c.Status(500).JSON(fiber.Map{"error": true, "msg": errGenerateAccessJWT.Error()})
			return
		}

		// Create JWT refresh_token
		refreshToken, errGenerateRefreshJWT := utils.GenerateRefreshJWT(accessToken)
		if errGenerateRefreshJWT != nil {
			// Fail create JWT token
			c.Status(500).JSON(fiber.Map{"error": true, "msg": errGenerateRefreshJWT.Error()})
			return
		}

		c.JSON(fiber.Map{
			"error": false,
			"msg":   nil,
			"jwt": fiber.Map{
				"access_token":  accessToken,
				"refresh_token": refreshToken,
				"expires_in":    time.Now().Add(72 * time.Hour).Unix(),
			},
		})
	} else {
		// Fail authentication
		c.Status(401).JSON(fiber.Map{"error": true, "msg": "incorrect email or password"})
		return
	}
}
