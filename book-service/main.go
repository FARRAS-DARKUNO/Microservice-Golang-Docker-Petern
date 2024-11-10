// book-service/main.go
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/FARRAS-DARKUNO/library-management/book-service/config"
    "github.com/FARRAS-DARKUNO/library-management/book-service/routers"
)

func main() {
    app := fiber.New()

    // Inisialisasi database
    config.InitDB()
    
    routes.SetupBookRoutes(app)

    app.Listen(":3001")
}
