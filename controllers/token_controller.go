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
