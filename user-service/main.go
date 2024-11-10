package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/FARRAS-DARKUNO/library-management/user-service/config"
    "github.com/FARRAS-DARKUNO/library-management/user-service/routes"
)

func main() {
    app := fiber.New()

    config.InitDB()

    routes.AuthRoutes(app)

    app.Listen(":3004")
}