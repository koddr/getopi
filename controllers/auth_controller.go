package controllers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	"github.com/koddr/getopi/models"
	"github.com/koddr/getopi/stores"
	"github.com/koddr/getopi/utils"
	nanoid "github.com/matoous/go-nanoid"
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

		// Create new token
		newToken := &models.Token{
			ID:          uuid.New(),
			UserID:      user.ID,
			CreatedAt:   time.Now(),
			ExpiredAt:   time.Now().Add(72 * time.Hour), // 72 hours to expire
			AccessToken: accessToken,
		}

		// Create JWT token
		tokenData, errCreateToken := db.CreateToken(newToken)
		if errCreateToken != nil {
			// User not found
			c.Status(403).JSON(fiber.Map{"error": true, "msg": "token not created", "auth": nil})
			return
		}

		c.JSON(fiber.Map{"error": false, "msg": nil, "auth": tokenData})
	} else {
		// Fail authentication
		c.Status(401).JSON(fiber.Map{"error": true, "msg": "incorrect email or password"})
		return
	}
}

// RefreshToken ...
func RefreshToken(c *fiber.Ctx) {
	// Get data from JWT
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	// Check UUID from current user
	currentUserID, errParse := uuid.Parse(claims["id"].(string))
	if errParse != nil {
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errParse.Error()})
		return
	}

	// Struct for arrived from frontend JWT token
	arrivedToken := &models.Token{}

	// Check received JSON data
	if errBodyParser := c.BodyParser(arrivedToken); errBodyParser != nil {
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

	// Check, if arrived refresh token is exists
	storedToken, errFindTokenByID := db.FindTokenByID(arrivedToken.ID)
	if errFindTokenByID != nil {
		// User not found
		c.Status(404).JSON(fiber.Map{"error": true, "msg": "token not found"})
		return
	}

	// Only equal JWT refresh_token can be refreshed
	if arrivedToken.ID == storedToken.ID {
		// Create new JWT access_token
		accessToken, errGenerateAccessJWT := utils.GenerateAccessJWT("user", claims["id"].(string))
		if errGenerateAccessJWT != nil {
			// Fail create JWT token
			c.Status(500).JSON(fiber.Map{"error": true, "msg": errGenerateAccessJWT.Error()})
			return
		}

		// Create new JWT token data
		newToken := &models.Token{
			ID:          uuid.New(),
			UserID:      currentUserID,
			CreatedAt:   time.Now(),
			ExpiredAt:   time.Now().Add(72 * time.Hour), // 72 hours to expire
			AccessToken: accessToken,
		}

		// Create new JWT token
		errRefreshTokenByID := db.RefreshTokenByID(storedToken.ID, newToken)
		if errRefreshTokenByID != nil {
			// Fail create new JWT token
			c.Status(500).JSON(fiber.Map{"error": true, "msg": "token not refreshed", "auth": nil})
			return
		}

		// Return new JWT token data to frontend in JSON format
		c.JSON(fiber.Map{"error": false, "msg": nil, "auth": newToken})
	} else {
		// Fail refresh JWT token
		c.Status(403).JSON(fiber.Map{"error": true, "msg": "permission denied", "user": nil})
		return
	}
}

// ForgetPassword email string..error.
func ForgetPassword(c *fiber.Ctx) {
	// Struct for restore password by email
	type forgetData struct {
		Email string `json:"email" validate:"required,email"`
	}

	// Create new forget password struct
	forget := &forgetData{}

	// Create new validator
	validate := utils.Validate("forget-password")

	// Check received JSON data
	if errBodyParser := c.BodyParser(forget); errBodyParser != nil {
		// Incorrect data
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errBodyParser.Error()})
		return
	}

	// Check fields validation
	if errValidate := validate.Struct(forget); errValidate != nil {
		// Return invalid fields
		c.Status(500).JSON(fiber.Map{"error": true, "msg": utils.ValidateErrors(errValidate)})
		return
	}

	// Create DB connection
	db, errConnectDB := stores.OpenStore()
	if errConnectDB != nil {
		// Fail DB connection
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errConnectDB.Error()})
		return
	}

	// Find user by email
	_, errFindUserByEmail := db.FindUserByEmail(forget.Email)
	if errFindUserByEmail != nil {
		// User not found
		c.Status(404).JSON(fiber.Map{"error": true, "msg": "user not found", "user": nil})
		return
	}

	// Create restore code
	restoreCode, errRestoreCode := nanoid.Generate("123456abcdef", 6)
	if errRestoreCode != nil {
		// Fail create restore code
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errRestoreCode.Error()})
		return
	}

	// Create new email sender
	sender := utils.NewEmailSender(
		utils.GetDotEnvValue("SERVER_EMAIL"),
		utils.GetDotEnvValue("SERVER_EMAIL_PASSWORD"),
		utils.GetDotEnvValue("SMTP_SERVER"),
		utils.GetDotEnvValue("SMTP_PORT"),
	)

	// Send email process
	if errSendHTMLEmail := sender.SendHTMLEmail(
		"templates/email-forgot-password.html",
		[]string{forget.Email},
		"Your restore code",
		fiber.Map{"code": restoreCode},
	); errSendHTMLEmail != nil {
		// Fail send restore code to email
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errSendHTMLEmail.Error()})
		return
	}

	c.JSON(fiber.Map{"error": false, "msg": nil, "restore_code": restoreCode})
}
