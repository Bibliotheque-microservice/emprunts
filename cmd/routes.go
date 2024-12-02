package main

import (
	"github.com/Bibliotheque-microservice/emprunts/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)

	// Route pour mettre à jour l'emprunt
	app.Put("/", handlers.UpdateEmprunts)

	// Nouvelle route pour créer un emprunt
	app.Post("/emprunt", handlers.CreateEmprunt)
}
