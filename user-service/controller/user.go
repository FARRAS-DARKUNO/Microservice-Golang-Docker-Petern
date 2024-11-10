package controllers

import (
	"os"
	"time"

	"github.com/FARRAS-DARKUNO/library-management/user-service/config"
	"github.com/FARRAS-DARKUNO/library-management/user-service/models"
	"github.com/gofiber/fiber/v2"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Login function to authenticate user
func Login(c *fiber.Ctx) error {
	var input models.User
	var user models.User

	// Parse the request body
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing request body")
	}

	// Check if user exists in database
	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid username or password")
	}

	// Compare the passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid username or password")
	}

	// Create JWT token
	token, err := generateJWT(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error generating token")
	}

	// Return the response with the token
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
	})
}

// Register function to create a new user
func Register(c *fiber.Ctx) error {
	var input models.User
	var user models.User

	// Parse the request body
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing request body")
	}

	// Check if user already exists
	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err == nil {
		return c.Status(fiber.StatusConflict).SendString("Username already exists")
	}

	// Encrypt the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error hashing password")
	}
	input.Password = string(hashedPassword)

	// Create the user in the database
	if err := config.DB.Create(&input).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating user")
	}

	// Return the response with the token
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Registration successful",
	})
}

// Function to generate JWT token
func generateJWT(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":       user.ID,
		"role":      user.Role,
		"role_code": user.RoleCode,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // token valid for 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}
