package main

import (
	"github.com/Bibliotheque-microservice/emprunts/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
}
