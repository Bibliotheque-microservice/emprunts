package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Bibliotheque-microservice/emprunts/cron"
	"github.com/Bibliotheque-microservice/emprunts/database"
	rabbitmq "github.com/Bibliotheque-microservice/emprunts/rabbitmq"
	"github.com/Bibliotheque-microservice/emprunts/structures"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()

	rabbitmq.InitRabbitMQ()
	defer rabbitmq.CloseRabbitMQ()

	go func() {
		penality_msgs := rabbitmq.ConsumeMessages("user_penalties_queue")

		for msg := range penality_msgs {
			switch msg.RoutingKey {
			case "user.v1.penalities.new":
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
	}()

	go func() {
		emprunts_msg := rabbitmq.ConsumeMessages("emprunts_finished_queue")

		for msg := range emprunts_msg {
			switch msg.RoutingKey {
			case "emprunts.v1.finished":
				var jsonData interface{}
				err := json.Unmarshal(msg.Body, &jsonData)
				if err != nil {
					log.Printf("Erreur de parsing JSON : %v", err)
					continue
				}
				log.Printf("Message JSON reçu : %+v", jsonData)
			default:
				log.Printf("Message non géré avec Routing Key : %s", msg.RoutingKey)
				log.Printf("Contenu brut du message : %s", string(msg.Body))
			}
		}
	}()

	go func() {
		cron.StartCron()
	}()
	app := fiber.New()
	setupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Listen(fmt.Sprintf(":%s", port))
}
