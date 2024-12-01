package main

import (
	"encoding/json"
	"log"

	"github.com/Bibliotheque-microservice/emprunts/database"
	rabbitmq "github.com/Bibliotheque-microservice/emprunts/rabbitMQ"
	"github.com/Bibliotheque-microservice/emprunts/structures"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()

	rabbitmq.InitRabbitMQ()
	defer rabbitmq.CloseRabbitMQ()

	penality_msgs := rabbitmq.ConsumeMessages("user_penalties_queue")

	for msg := range penality_msgs {
		switch msg.RoutingKey {
		case "user.v1.penalities.new":

			// Parsing JSON and assign to jsonData
			var jsonData structures.PenaltyMessage
			err := json.Unmarshal(msg.Body, &jsonData)
			if err != nil {
				log.Printf("Erreur de parsing JSON : %v", err)
				continue
			}
			log.Printf("Message JSON reçu : %+v", jsonData)
			log.Print(jsonData.Amount)

		case "user.v1.penalities.paid":
			log.Printf("Message reçu : %s", string(msg.Body))
		default:
			log.Printf("Message non géré avec Routing Key : %s", msg.RoutingKey)
			log.Printf("Contenu brut du message : %s", string(msg.Body))
		}

	}
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":5000")
}
