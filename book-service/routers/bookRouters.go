package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/FARRAS-DARKUNO/library-management/book-service/controllers"
)

// SetupBookRoutes mendefinisikan semua routing yang ada
func SetupBookRoutes(app *fiber.App) {
	app.Post("/books", controllers.CreateBook)
	app.Get("/books", controllers.GetBooks)
	app.Get("/books/:id", controllers.GetBookByID)
	app.Put("/books/:id", controllers.UpdateBook)
	app.Delete("/books/:id", controllers.DeleteBook)
	
	app.Post("/bookstock", controllers.CreateBookStock)
	app.Post("/borrow", controllers.CreateBorrowBook)

	// Routing untuk pengembalian buku
	app.Put("/return/:id", controllers.ReturnBook) // Menambahkan endpoint untuk pengembalian buku
}
