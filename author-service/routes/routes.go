package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/FARRAS-DARKUNO/library-management/author-service/controllers"
)

// SetupAuthorRoutes mendefinisikan semua routing untuk AuthorService
func SetupAuthorRoutes(app *fiber.App) {
	app.Post("/authors", controllers.CreateAuthor)
	app.Get("/authors", controllers.GetAuthors)
	app.Get("/authors/:id", controllers.GetAuthorByID)
	app.Put("/authors/:id", controllers.UpdateAuthor)
	app.Delete("/authors/:id", controllers.DeleteAuthor)
}
