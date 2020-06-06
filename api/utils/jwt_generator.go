package utils

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

// GenerateAccessJWT ...
func GenerateAccessJWT(permission, id string) (string, error) {
	// Get secret JWT token from .env
	secretToken := os.Getenv("JWT_SECRET_TOKEN")

	// Create JWT token and claims
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	// Set superadmin claims
	switch permission {
	case "admin":
		claims["is_admin"] = true
	default:
		claims["is_admin"] = false
	}

	// Set public claims
	claims["id"] = id

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(secretToken))
	if err != nil {
		return "", err
	}

	return t, nil
}
