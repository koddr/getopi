package controllers

import (
	"time"

	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	"github.com/koddr/getopi/models"
	"github.com/koddr/getopi/stores"
	"github.com/koddr/getopi/utils"
	nanoid "github.com/matoous/go-nanoid"
)

// Authentication ...
func Authentication(c *fiber.Ctx) {
	// Create new User struct
	authData := &models.Auth{}

	// Check received JSON data
	if errBodyParser := c.BodyParser(authData); errBodyParser != nil {
		// Incorrect data
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errBodyParser.Error()})
		return
	}

	// Create new validator
	validate := utils.Validate("auth")

	// Check fields validation
	if errValidate := validate.Struct(authData); errValidate != nil {
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
	user, errFindUserByEmail := db.FindUserByEmail(authData.Email)
	if errFindUserByEmail != nil {
		// User not found
		c.Status(404).JSON(fiber.Map{"error": true, "msg": "user not found", "user": nil})
		return
	}

	// Check password
	if utils.ComparePasswords(user.PasswordHash, authData.Password) {
		// Create JWT access_token
		accessToken, errGenerateAccessJWT := utils.GenerateAccessJWT("user", user.ID.String())
		if errGenerateAccessJWT != nil {
			// Fail create JWT token
			c.Status(500).JSON(fiber.Map{"error": true, "msg": errGenerateAccessJWT.Error()})
			return
		}

		// Create new token data
		newTokenData := &models.Token{
			ID:          uuid.New(),
			UserID:      user.ID,
			CreatedAt:   time.Now(),
			ExpiredAt:   time.Now().Add(72 * time.Hour), // 72 hours to expire
			AccessToken: accessToken,
		}

		// Delete exists JWT token
		errDeleteTokenByUserID := db.DeleteTokenByUserID(user.ID)
		if errDeleteTokenByUserID != nil {
			// Fail delete exists JWT token
			c.Status(500).JSON(fiber.Map{"error": true, "msg": "token not deleted", "auth": nil})
			return
		}

		// Create JWT token
		newToken, errCreateToken := db.CreateToken(newTokenData)
		if errCreateToken != nil {
			// User not found
			c.Status(403).JSON(fiber.Map{"error": true, "msg": "token not created", "auth": nil})
			return
		}

		c.JSON(fiber.Map{"error": false, "msg": nil, "auth": newToken})
	} else {
		// Fail authentication
		c.Status(401).JSON(fiber.Map{"error": true, "msg": "incorrect email or password"})
		return
	}
}

// ForgetPasswordIssue ...
func ForgetPasswordIssue(c *fiber.Ctx) {
	// Create new forget password struct
	forgetData := &models.ForgetPassword{}

	// Check received JSON data
	if errBodyParser := c.BodyParser(forgetData); errBodyParser != nil {
		// Incorrect data
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errBodyParser.Error()})
		return
	}

	// Create new validator
	validate := utils.Validate("forget-password")

	// Check fields validation
	if errValidate := validate.Struct(forgetData); errValidate != nil {
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
	user, errFindUserByEmail := db.FindUserByEmail(forgetData.Email)
	if errFindUserByEmail != nil {
		// User not found
		c.Status(404).JSON(fiber.Map{"error": true, "msg": "user not found"})
		return
	}

	// Create reset code
	resetCode, errResetCode := nanoid.Generate("1234567890abcdefxyz", 6)
	if errResetCode != nil {
		// Fail create restore code
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errResetCode.Error()})
		return
	}

	// Create new reset code issue
	newResetCode := &models.ResetCode{
		ID:        uuid.New(),
		UserID:    user.ID,
		ResetCode: resetCode,
	}

	// Delete reset issue by code
	errDeleteResetPasswordIssueByUserID := db.DeleteResetPasswordIssueByUserID(user.ID)
	if errDeleteResetPasswordIssueByUserID != nil {
		// Fail delete reset password code
		c.Status(500).JSON(fiber.Map{"error": true, "msg": "reset code not deleted"})
		return
	}

	// Create new reset code issue
	errCreateResetPasswordIssue := db.CreateResetPasswordIssue(newResetCode)
	if errCreateResetPasswordIssue != nil {
		// Fail create new issue
		c.Status(500).JSON(fiber.Map{"error": true, "msg": "reset code not created"})
		return
	}

	// Create new email sender
	sender := utils.NewEmailSender(
		utils.GetDotEnvValue("SERVER_EMAIL"),
		utils.GetDotEnvValue("SERVER_EMAIL_PASSWORD"),
		utils.GetDotEnvValue("SMTP_SERVER"),
		utils.GetDotEnvValue("SMTP_PORT"),
	)

	// Send email with password reset link
	if errSendHTMLEmail := sender.SendHTMLEmail(
		"templates/email-forgot-password.html", []string{forgetData.Email},
		"Your password reset code", fiber.Map{"code": resetCode},
	); errSendHTMLEmail != nil {
		// Fail send restore code to email
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errSendHTMLEmail.Error()})
		return
	}

	c.JSON(fiber.Map{"error": false, "msg": nil})
}

// ForgetPasswordCheckResetCode ...
func ForgetPasswordCheckResetCode(c *fiber.Ctx) {
	// Create new reset code struct
	resetCodeData := &models.ResetCode{}

	// Check received JSON data
	if errBodyParser := c.BodyParser(resetCodeData); errBodyParser != nil {
		// Incorrect data
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errBodyParser.Error()})
		return
	}

	// Create new validator
	validate := utils.Validate("reset-code")

	// Check fields validation
	if errValidate := validate.Struct(resetCodeData); errValidate != nil {
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

	// Find reset issue by code
	errFindResetPasswordIssueByCode := db.FindResetPasswordIssueByCode(resetCodeData.ResetCode)
	if errFindResetPasswordIssueByCode != nil {
		// Reset code not found
		c.Status(404).JSON(fiber.Map{"error": true, "msg": "reset code not found"})
		return
	}

	// Delete reset issue by code
	errDeleteResetPasswordIssueByCode := db.DeleteResetPasswordIssueByCode(resetCodeData.ResetCode)
	if errDeleteResetPasswordIssueByCode != nil {
		// User not found
		c.Status(500).JSON(fiber.Map{"error": true, "msg": "reset code not deleted"})
		return
	}

	c.JSON(fiber.Map{"error": false, "msg": nil})
}
