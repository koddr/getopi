package controllers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	"github.com/koddr/getopi/models"
	"github.com/koddr/getopi/stores"
	"github.com/koddr/getopi/utils"
)

// ShowUserByUsername ...
//
// TODO: Add description
//
func ShowUserByUsername(c *fiber.Ctx) {
	db, err := stores.OpenStore()
	if err != nil {
		// DB connection error
		c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
		return
	}

	// Get username from URL
	username := c.Params("username")

	// Find user by username
	user, err := db.FindUserByUsername(username)
	if err != nil {
		// User not found
		c.Status(404).JSON(fiber.Map{"error": true, "msg": err.Error(), "user": nil})
		return
	}

	// Hide PasswordHash field from JSON output
	user.PasswordHash = ""

	// OK result
	c.JSON(fiber.Map{"error": false, "msg": nil, "user": user})
}

// ShowUsers ...
//
// TODO: Add description
//
func ShowUsers(c *fiber.Ctx) {
	db, err := stores.OpenStore()
	if err != nil {
		// DB connection error
		c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
		return
	}

	// Select all users
	users, err := db.GetUsers()
	if err != nil {
		// Users not found
		c.Status(404).JSON(fiber.Map{"error": true, "msg": err.Error(), "count": 0, "users": nil})
		return
	}

	// Hide PasswordHash field from JSON output
	for index := range users {
		users[index].PasswordHash = ""
	}

	// OK result
	c.JSON(fiber.Map{"error": false, "msg": nil, "count": len(users), "users": users})
}

// CreateUser ...
//
// TODO: Add description
//
func CreateUser(c *fiber.Ctx) {
	db, err := stores.OpenStore()
	if err != nil {
		// DB connection error
		c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
		return
	}

	// Create new User struct
	user := &models.User{}

	// Check received JSON data
	if err := c.BodyParser(user); err != nil {
		// Incorrect data
		c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
		return
	}

	// Validate data
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		errorFields := map[string]interface{}{}
		for _, err := range err.(validator.ValidationErrors) {
			errorFields[err.Field()] = "need " + err.Type().Name() + " got " + err.Kind().String()
		}
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errorFields})
		return
	}

	// Set init user data
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Time{}
	user.Username = user.ID.String()[:13] // first 13 letters of ID
	user.PasswordHash = utils.GeneratePassword(user.PasswordHash)
	user.UserStatus = 1
	user.UserAttrs = models.UserAttrs{}

	// Create new user
	if err := db.CreateUser(user); err != nil {
		// Not inserted new user to DB
		c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
		return
	}

	// Hide PasswordHash field from JSON output
	user.PasswordHash = ""

	// OK result
	c.JSON(fiber.Map{"error": false, "msg": nil, "user": user})
}

// UpdateUser ...
//
// TODO: Add description
//
func UpdateUser(c *fiber.Ctx) {
	// Get data from JWT
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	// Check UUID from current user
	currentUserID, err := uuid.Parse(claims["id"].(string))
	if err != nil {
		c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
		return
	}

	// Create DB connection
	db, err := stores.OpenStore()
	if err != nil {
		// DB connection error
		c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
		return
	}

	// Create new User struct
	user := &models.User{}

	// Check received JSON data
	if err := c.BodyParser(user); err != nil {
		// Incorrect data
		c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
		return
	}

	// Check if user with given Username is exists
	if _, err := db.FindUserByID(user.ID); err != nil {
		// User not found
		c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
		return
	}

	// Only owner can update itself
	if currentUserID == user.ID {
		// Set user data to update
		user.UpdatedAt = time.Now()

		// Update user
		if err := db.UpdateUser(user); err != nil {
			// Not inserted new user to DB
			c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
			return
		}

		// Hide PasswordHash field from JSON output
		user.PasswordHash = ""

		// OK result
		c.JSON(fiber.Map{"error": false, "msg": nil, "user": user})
	} else {
		// If it's not owner
		c.Status(500).JSON(fiber.Map{"error": true, "msg": "permission denied", "user": nil})
		return
	}
}

// DeleteUser ...
//
// TODO: Add description
//
func DeleteUser(c *fiber.Ctx) {
	// Get data from JWT
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	isAdmin := claims["is_admin"].(bool)

	// Check, if current user request's from admin
	if isAdmin {
		db, err := stores.OpenStore()
		if err != nil {
			// DB connection error
			c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
			return
		}

		// Create new User struct
		user := &models.User{}

		// Check received JSON data
		if err := c.BodyParser(user); err != nil {
			// Incorrect data
			c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
			return
		}

		// Check ID (UUID) for empty value
		if user.ID == uuid.Nil {
			// User not found
			c.Status(500).JSON(fiber.Map{"error": true, "msg": "incorrect ID"})
			return
		}

		// Deleter user
		if err := db.DeleteUser(user.ID); err != nil {
			// Not inserted new user to DB
			c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
			return
		}

		// OK result
		c.JSON(fiber.Map{"error": false, "msg": nil})
	} else {
		// If it's not admin
		c.Status(500).JSON(fiber.Map{"error": true, "msg": "permission denied"})
		return
	}
}
