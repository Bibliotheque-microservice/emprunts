package main

import (
	"github.com/Bibliotheque-microservice/emprunts/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, ma vie!")
	})

	app.Listen(":5000")
}
