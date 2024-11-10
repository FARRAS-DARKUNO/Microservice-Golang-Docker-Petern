// book-service/main.go
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/FARRAS-DARKUNO/library-management/author-service/config"
    "github.com/FARRAS-DARKUNO/library-management/author-service/routes"
)

func main() {
    app := fiber.New()

    // Inisialisasi database
    config.InitDB()

    routes.SetupAuthorRoutes(app)

    app.Listen(":3002")
}
