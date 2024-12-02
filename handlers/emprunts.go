package handlers

import (
	"encoding/json"
	"log"
	"time"

	"github.com/Bibliotheque-microservice/emprunts/database"
	"github.com/Bibliotheque-microservice/emprunts/models"
	"github.com/Bibliotheque-microservice/emprunts/rabbitmq"
	"github.com/Bibliotheque-microservice/emprunts/services" // Import des services pour vérifier livre et utilisateur
	"github.com/Bibliotheque-microservice/emprunts/structures"
	"github.com/gofiber/fiber/v2"
)

// Route pour vérifier et créer un emprunt
func CreateEmprunt(c *fiber.Ctx) error {
	// Parse la requête JSON pour obtenir l'ID de l'utilisateur et du livre
	var request structures.EmpruntRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Vérifier la disponibilité du livre
	available, err := services.CheckBookAvailability(request.BookID)
	if err != nil || !available {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Le livre n'est pas disponible"})
	}

	// Vérifier l'état de l'utilisateur (actif et pas de pénalités)
	userAuthorized, err := services.CheckUserStatus(request.UserID)
	if err != nil || !userAuthorized {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Utilisateur non autorisé"})
	}

	// Créer l'emprunt dans la base de données
	emprunt := models.Emprunt{
		UtilisateurID:   request.UserID,
		LivreID:         request.BookID,
		DateEmprunt:     time.Now(),
		DateRetourPrevu: time.Now().Add(14 * 24 * time.Hour), // 2 semaines de durée
	}

	err = database.DB.Db.Create(&emprunt).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erreur lors de la création de l'emprunt"})
	}

	// Publier un message via RabbitMQ pour notifier l'emprunt
	empruntMessage := map[string]interface{}{
		"livreId":       request.BookID,
		"disponible":    false,
		"idUtilisateur": request.UserID,
	}

	// Convertir le message en JSON
	message, err := json.Marshal(empruntMessage)
	if err != nil {
		log.Printf("Erreur lors de la création du message RabbitMQ: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erreur de publication du message"})
	}

	// Publier le message à RabbitMQ
	rabbitmq.PublishMessage("emprunts_exchange", "emprunts.v1.created", string(message))

	// Répondre avec succès et autoriser l'emprunt
	return c.JSON(fiber.Map{"autorisé": true, "message": "Emprunt créé avec succès"})
}
