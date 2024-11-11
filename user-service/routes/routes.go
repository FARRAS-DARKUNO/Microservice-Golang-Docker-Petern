package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/FARRAS-DARKUNO/library-management/user-service/controller"
)

func AuthRoutes(app *fiber.App) {

	app.Post("/login", controllers.Login)

	app.Post("/register", controllers.Register)

	app.Get("/user/:id", controllers.GetUserByID)

	app.Use(app)
}