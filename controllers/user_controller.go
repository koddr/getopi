package controllers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	"github.com/koddr/getopi/models"
	"github.com/koddr/getopi/stores"
	"github.com/koddr/getopi/utils"
)

// UserController ...
//
// TODO: Add description
//
func UserController(c *fiber.Ctx) {
	db, err := stores.OpenStore()
	if err != nil {
		// DB connection error
		c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
		return
	}

	// Get username from URL
	username := c.Params("username")

	// Select user by username
	user, err := db.UserByUsername(username)
	if err != nil {
		// User not found
		c.Status(404).JSON(fiber.Map{"error": false, "msg": err.Error()})
		return
	}

	// Hide PasswordHash field from JSON output
	user.PasswordHash = ""

	// OK result
	c.JSON(fiber.Map{"error": false, "msg": "ok", "user": user})
}

// UsersController ...
//
// TODO: Add description
//
func UsersController(c *fiber.Ctx) {
	db, err := stores.OpenStore()
	if err != nil {
		// DB connection error
		c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
		return
	}

	// Select all users
	users, err := db.Users()
	if err != nil {
		// Users not found
		c.Status(404).JSON(fiber.Map{"error": false, "msg": err.Error()})
		return
	}

	// Hide PasswordHash field from JSON output
	for index := range users {
		users[index].PasswordHash = ""
	}

	// OK result
	c.JSON(fiber.Map{"error": false, "msg": "ok", "count": len(users), "users": users})
}

// UserCreateController ...
//
// TODO: Add description
//
func UserCreateController(c *fiber.Ctx) {
	db, err := stores.OpenStore()
	if err != nil {
		// DB connection error
		c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
		return
	}

	// Create new User struct
	user := new(models.User)

	// Check received JSON data
	if err := c.BodyParser(user); err != nil {
		// Incorrect data
		c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
		return
	}

	// Check Email for empty value
	if user.Email == "" || user.PasswordHash == "" {
		// Incorrect data
		c.Status(500).JSON(fiber.Map{"error": true, "msg": "incorrect Email or Password"})
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

	// OK result
	c.JSON(fiber.Map{"error": false, "msg": "ok"})
}

// UserUpdateController ...
//
// TODO: Add description
//
func UserUpdateController(c *fiber.Ctx) {
	// Get data from JWT
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	currentUserID, _ := uuid.Parse(claims["id"].(string))

	db, err := stores.OpenStore()
	if err != nil {
		// DB connection error
		c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
		return
	}

	// Create new User struct
	user := new(models.User)

	// Check received JSON data
	if err := c.BodyParser(user); err != nil {
		// Incorrect data
		c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
		return
	}

	// Check ID (UUID) and Username (string) for empty values
	if user.ID == uuid.Nil {
		// User not found
		c.Status(500).JSON(fiber.Map{"error": true, "msg": "incorrect ID"})
		return
	}

	// Check if user with given Username is exists
	if _, err := db.User(user.ID); err != nil {
		// User not found
		c.Status(404).JSON(fiber.Map{"error": false, "msg": err.Error()})
		return
	}

	if currentUserID == user.ID {
		// Set user data to update
		user.UpdatedAt = time.Now()

		// Update user
		if err := db.UpdateUser(user); err != nil {
			// Not inserted new user to DB
			c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
			return
		}

		// OK result
		c.JSON(fiber.Map{"error": false, "msg": "ok"})
	} else {
		// If it's not owner
		c.Status(500).JSON(fiber.Map{"error": true, "msg": "permission denied"})
		return
	}
}

// UserDeleteController ...
//
// TODO: Add description
//
func UserDeleteController(c *fiber.Ctx) {
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
		user := new(models.User)

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
		c.JSON(fiber.Map{"error": false, "msg": "ok"})
	} else {
		// If it's not admin
		c.Status(500).JSON(fiber.Map{"error": true, "msg": "permission denied"})
		return
	}
}
