package handlers

import (
	"time"

	"github.com/Bibliotheque-microservice/emprunts/database"
	"github.com/Bibliotheque-microservice/emprunts/models"
	rabbitmq "github.com/Bibliotheque-microservice/emprunts/rabbitMQ"
	"github.com/Bibliotheque-microservice/emprunts/structures"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello ma vie!")
}

func UpdateEmprunts(c *fiber.Ctx) error {
	// Parse la requete
	var empruntRequest structures.EmpruntReturned
	if err := c.BodyParser(&empruntRequest); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Si l'emprunt a été retourné, updater la date de retour
	if empruntRequest.Returned {
		err := database.DB.Db.Model(&models.Emprunt{}).
			Where("id_emprunt = ?", empruntRequest.EmpruntID).
			Update("date_retour_effectif", time.Now()).Error

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Erreur lors de la mise à jour de l'emprunt avec bdd.",
				"error":   err.Error(),
			})
		}

		// retrieve the whole emprunt
		var updatedEmprunt models.Emprunt
		err = database.DB.Db.First(&updatedEmprunt, "id_emprunt = ?", empruntRequest.EmpruntID).Error

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Erreur lors de la mise à jour de l'emprunt avec bdd.",
				"error":   err.Error(),
			})
		}

		rabbitmq.PublishMessage("emprunts_exchange", "emprunts.v1.finished", updatedEmprunt)

		// Envoyer un message aux cosnumers pour indiquer que le livre a été retourné
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Date de retour mise à jour avec succès.",
		})
	} else {
		return c.Status(400).SendString("Livre not returned")
	}

}
