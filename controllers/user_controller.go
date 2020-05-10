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

// ShowUserByUsername show exists user by username or 404 error
func ShowUserByUsername(c *fiber.Ctx) {
	// Create DB connection
	db, errConnectDB := stores.OpenStore()
	if errConnectDB != nil {
		// DB connection error
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errConnectDB.Error()})
		return
	}

	// Get username from URL
	username := c.Params("username")

	// Find user by username
	user, errFindUserByUsername := db.FindUserByUsername(username)
	if errFindUserByUsername != nil {
		// User not found
		c.Status(404).JSON(fiber.Map{"error": true, "msg": errFindUserByUsername.Error(), "user": nil})
		return
	}

	// Hide PasswordHash field from JSON output
	user.PasswordHash = ""

	// OK result
	c.JSON(fiber.Map{"error": false, "msg": nil, "user": user})
}

// ShowUsers show all exists users or 404 error
func ShowUsers(c *fiber.Ctx) {
	// Create DB connection
	db, errConnectDB := stores.OpenStore()
	if errConnectDB != nil {
		// DB connection error
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errConnectDB.Error()})
		return
	}

	// Select all users
	users, errGetUsers := db.GetUsers()
	if errGetUsers != nil {
		// Users not found
		c.Status(404).JSON(fiber.Map{"error": true, "msg": errGetUsers.Error(), "count": 0, "users": nil})
		return
	}

	// Hide PasswordHash field from JSON output
	for index := range users {
		users[index].PasswordHash = ""
	}

	// OK result
	c.JSON(fiber.Map{"error": false, "msg": nil, "count": len(users), "users": users})
}

// CreateUser create new DB connection, gets JSON from request body, validate data,
// set init user values, clear password hash and create new user
func CreateUser(c *fiber.Ctx) {
	// Create DB connection
	db, errConnectDB := stores.OpenStore()
	if errConnectDB != nil {
		// Show DB connection error
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errConnectDB.Error()})
		return
	}

	// Create new validator
	validate := utils.Validate("user")

	// Create new User struct
	user := &models.User{}

	// Check received data from JSON body
	if errBodyParser := c.BodyParser(user); errBodyParser != nil {
		// Show incorrect JSON data
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errBodyParser.Error()})
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

	// Check fields validation
	if errValidate := validate.Struct(user); errValidate != nil {
		// Return invalid fields
		c.Status(500).JSON(fiber.Map{"error": true, "msg": utils.ValidateErrors(errValidate)})
		return
	}

	// Create new user with validated data
	if errCreateUser := db.CreateUser(user); errCreateUser != nil {
		// Show insert new row error
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errCreateUser.Error()})
		return
	}

	// Hide PasswordHash field from JSON output
	user.PasswordHash = ""

	// OK result
	c.Status(201).JSON(fiber.Map{"error": false, "msg": nil, "user": user})
}

// UpdateUser update user (by only its owner) by ID or 500 error
func UpdateUser(c *fiber.Ctx) {
	// Get data from JWT
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	// Check admin status
	isAdmin := claims["is_admin"].(bool)

	// Check UUID from current user
	currentUserID, errParse := uuid.Parse(claims["id"].(string))
	if errParse != nil {
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errParse.Error()})
		return
	}

	// Create DB connection
	db, errConnectDB := stores.OpenStore()
	if errConnectDB != nil {
		// DB connection error
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errConnectDB.Error()})
		return
	}

	// Create new validator
	validate := utils.Validate("user")

	// Create new User struct
	user := &models.User{}

	// Check received JSON data
	if errBodyParser := c.BodyParser(user); errBodyParser != nil {
		// Incorrect data
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errBodyParser.Error()})
		return
	}

	// Only if owner can update itself or it's admin
	if currentUserID == user.ID || isAdmin {

		// Check fields validation
		if errValidate := validate.Struct(user); errValidate != nil {
			// Return invalid fields
			c.Status(500).JSON(fiber.Map{"error": true, "msg": utils.ValidateErrors(errValidate)})
			return
		}

		// Check if user with given Username is exists
		if _, errFindUserByID := db.FindUserByID(user.ID); errFindUserByID != nil {
			// User not found
			c.Status(404).JSON(fiber.Map{"error": true, "msg": errFindUserByID.Error()})
			return
		}

		// Set user data to update
		user.UpdatedAt = time.Now()

		// Update user
		if errUpdateUser := db.UpdateUser(user); errUpdateUser != nil {
			// Fail update user in DB
			c.Status(500).JSON(fiber.Map{"error": true, "msg": errUpdateUser.Error()})
			return
		}

		// Hide PasswordHash field from JSON output
		user.PasswordHash = ""

		// OK result
		c.Status(202).JSON(fiber.Map{"error": false, "msg": nil, "user": user})
	} else {
		// If it's not owner
		c.Status(403).JSON(fiber.Map{"error": true, "msg": "permission denied", "user": nil})
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
		db, errConnectDB := stores.OpenStore()
		if errConnectDB != nil {
			// DB connection error
			c.Status(500).JSON(fiber.Map{"error": true, "msg": errConnectDB.Error()})
			return
		}

		// Create new User struct
		user := &models.User{}

		// Check received JSON data
		if errBodyParser := c.BodyParser(user); errBodyParser != nil {
			// Incorrect data
			c.Status(500).JSON(fiber.Map{"error": true, "msg": errBodyParser.Error()})
			return
		}

		// Deleter user
		if errDeleteUser := db.DeleteUser(user.ID); errDeleteUser != nil {
			// Not inserted new user to DB
			c.Status(500).JSON(fiber.Map{"error": true, "msg": errDeleteUser.Error()})
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
