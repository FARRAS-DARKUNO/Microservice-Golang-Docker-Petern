package main

import (
	"github.com/FARRAS-DARKUNO/library-management/category-service/config"
	"github.com/FARRAS-DARKUNO/library-management/category-service/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    config.InitDB()

    routes.CategoryRoutes(app)

    app.Listen(":3003")
}
