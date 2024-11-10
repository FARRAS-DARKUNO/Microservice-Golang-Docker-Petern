package middleware

import (
	"os"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var jwtSecretKey = os.Getenv("JWT_SECRET_KEY")

func IsAuthenticated(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Missing token")
	}

	// Remove Bearer prefix if exists
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
	}

	// Proceed to the next handler
	return c.Next()
}
