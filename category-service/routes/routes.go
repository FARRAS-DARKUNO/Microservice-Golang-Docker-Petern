package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/FARRAS-DARKUNO/library-management/category-service/controller"
)

func CategoryRoutes(app *fiber.App) {
	app.Post("/categories", controllers.CreateCategory)
	app.Get("/categories", controllers.GetCategories)
	app.Get("/categories/:id", controllers.GetCategoryByID)
	app.Put("/categories/:id", controllers.UpdateCategory)
	app.Delete("/categories/:id", controllers.DeleteCategory)
}
