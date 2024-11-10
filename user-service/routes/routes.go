package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/FARRAS-DARKUNO/library-management/user-service/controller"
	"github.com/FARRAS-DARKUNO/library-management/user-service/middleware"
)

func AuthRoutes(app *fiber.App) {
	// Login route
	app.Post("/login", controllers.Login)

	// Register route
	app.Post("/register", controllers.Register)

	// Protect other routes with JWT middleware
	app.Use(middleware.IsAuthenticated)
}