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

// ShowUserByUsername ...
// show exists user by username or 404 error
func ShowUserByUsername(c *fiber.Ctx) {
	// Get username from URL
	username := c.Params("username")

	// Create DB connection
	db, errConnectDB := stores.OpenStore()
	if errConnectDB != nil {
		// DB connection error
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errConnectDB.Error()})
		return
	}

	// Find user by username
	user, errFindUserByUsername := db.FindUserByUsername(username)
	if errFindUserByUsername != nil {
		// User not found
		c.Status(404).JSON(fiber.Map{"error": true, "msg": "user not found", "user": nil})
		return
	}

	// Hide PasswordHash field from JSON output
	user.PasswordHash = ""

	c.JSON(fiber.Map{"error": false, "msg": nil, "user": user})
}

// ShowUsers ...
// show all exists users or 404 error
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
		c.Status(404).JSON(fiber.Map{"error": true, "msg": "users not found", "count": 0, "users": nil})
		return
	}

	// Hide PasswordHash field from JSON output
	for index := range users {
		users[index].PasswordHash = ""
	}

	c.JSON(fiber.Map{"error": false, "msg": nil, "count": len(users), "users": users})
}

// CreateUser ...
// create new DB connection, gets JSON from request body, validate data,
// set init user values, clear password hash and create new user
func CreateUser(c *fiber.Ctx) {
	// Create new User struct
	user := &models.User{}

	// Create new validator
	validate := utils.Validate("user")

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

	// Create DB connection
	db, errConnectDB := stores.OpenStore()
	if errConnectDB != nil {
		// DB connection error
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errConnectDB.Error()})
		return
	}

	// Create new user with validated data
	if errCreateUser := db.CreateUser(user); errCreateUser != nil {
		// Fail create user
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errCreateUser.Error()})
		return
	}

	// Hide PasswordHash field from JSON output
	user.PasswordHash = ""

	c.Status(201).JSON(fiber.Map{"error": false, "msg": nil, "user": user})
}

// UpdateUser ...
// update user (by only its owner or admin) by ID or 500 error
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

	// Create new User struct
	user := &models.User{}

	// Check received JSON data
	if errBodyParser := c.BodyParser(user); errBodyParser != nil {
		// Fail parse JSON data
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errBodyParser.Error()})
		return
	}

	// Create DB connection
	db, errConnectDB := stores.OpenStore()
	if errConnectDB != nil {
		// DB connection error
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errConnectDB.Error()})
		return
	}

	// Check if user with given ID is exists
	if _, errFindUserByID := db.FindUserByID(user.ID); errFindUserByID != nil {
		// User not found
		c.Status(404).JSON(fiber.Map{"error": true, "msg": "user not found"})
		return
	}

	// Only if owner can update itself or it's admin
	if currentUserID == user.ID || isAdmin {
		// Create new validator
		validate := utils.Validate("user")

		// Check fields validation
		if errValidate := validate.Struct(user); errValidate != nil {
			// Return invalid fields
			c.Status(500).JSON(fiber.Map{"error": true, "msg": utils.ValidateErrors(errValidate)})
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

		c.Status(202).JSON(fiber.Map{"error": false, "msg": nil, "user": user})
	} else {
		// Fail update user
		c.Status(403).JSON(fiber.Map{"error": true, "msg": "permission denied", "user": nil})
		return
	}
}

// UpdateUserPassword ...
// update user password (by only its owner or admin) by ID or 500 error
func UpdateUserPassword(c *fiber.Ctx) {
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

	// Create new User struct
	user := &models.User{}

	// Check received JSON data
	if errBodyParser := c.BodyParser(user); errBodyParser != nil {
		// Incorrect data
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errBodyParser.Error()})
		return
	}

	// Create DB connection
	db, errConnectDB := stores.OpenStore()
	if errConnectDB != nil {
		// DB connection error
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errConnectDB.Error()})
		return
	}

	// Check if user with given ID is exists
	if _, errFindUserByID := db.FindUserByID(user.ID); errFindUserByID != nil {
		// User not found
		c.Status(404).JSON(fiber.Map{"error": true, "msg": "user not found"})
		return
	}

	// Only if owner can update itself or it's admin
	if currentUserID == user.ID || isAdmin {
		// Create new validator
		validate := utils.Validate("password")

		// Check fields validation
		if errValidate := validate.Struct(user); errValidate != nil {
			// Return invalid fields
			c.Status(500).JSON(fiber.Map{"error": true, "msg": utils.ValidateErrors(errValidate)})
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

		c.Status(202).JSON(fiber.Map{"error": false, "msg": nil, "user": user})
	} else {
		// Fail permission denied
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

	// Check admin status
	isAdmin := claims["is_admin"].(bool)

	// Check, if current user request's from admin
	if isAdmin {
		// Create DB connection
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
			// Fail parse JSON
			c.Status(500).JSON(fiber.Map{"error": true, "msg": errBodyParser.Error()})
			return
		}

		// Delete user
		if errDeleteUser := db.DeleteUser(user.ID); errDeleteUser != nil {
			// Fail delete user
			c.Status(500).JSON(fiber.Map{"error": true, "msg": errDeleteUser.Error()})
			return
		}

		c.JSON(fiber.Map{"error": false, "msg": nil})
	} else {
		// Fail permission denied
		c.Status(500).JSON(fiber.Map{"error": true, "msg": "permission denied"})
		return
	}
}
