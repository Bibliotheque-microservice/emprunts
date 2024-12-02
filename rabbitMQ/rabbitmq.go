package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Bibliotheque-microservice/emprunts/myutils"
	"github.com/rabbitmq/amqp091-go"
)

var conn *amqp091.Connection
var ch *amqp091.Channel

func InitRabbitMQ() {
	var err error

	url := fmt.Sprintf(
		"amqp://%s:%s@rabbitmq-emprunts:5672/",
		os.Getenv("BROKER_USER"),
		os.Getenv("BROKER_PASSWORD"),
	)

	// Connect channel
	conn, err = amqp091.Dial(url)
	myutils.FailOnError(err, "Failed to connect to RabbitMQ")

	// Open channel
	ch, err = conn.Channel()
	myutils.FailOnError(err, "Failed to open a channel")

	// Declare exchange
	err = ch.ExchangeDeclare(
		"emprunts_exchange",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	myutils.FailOnError(err, "Failed to declare the exchange")

	err = ch.ExchangeDeclare(
		"penality_exchange",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	myutils.FailOnError(err, "Failed to declare the exchange")

	// Declare queues

	// Used to alert book not available anymore
	DeclareAndBindQueue("emprunts_exchange", "emprunts_created_queue", "emprunts.v1.created")
	// Used to alert book reserved
	DeclareAndBindQueue("emprunts_exchange", "emprunts_finished_queue", "emprunts.v1.finished")

	// Used to alert penality TO BE REMOVED
	DeclareAndBindQueue("penality_exchange", "user_penalties_queue", "user.v1.penalities.paid")

	// Used to alert user that for a new penality
	DeclareAndBindQueue("penality_exchange", "user_penalties_queue", "user.v1.penalities.new")
}

// PublishMessage publishes a message to a specific routing key
func PublishMessage(exchangeName string, routingKey string, message interface{}) {

	messageJSON, err := json.Marshal(message)
	myutils.FailOnError(err, "Failed to marshal message to JSON")

	err = ch.Publish(
		exchangeName,
		routingKey, // routing key utile pour consume que cette route
		false,
		false, // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        messageJSON,
		})
	myutils.FailOnError(err, "Failed to publish a message")
	log.Printf("Published message to %s: %s", routingKey, message)
}

// ConsumeMessages starts consuming messages from a specific queue
func ConsumeMessages(queueName string) <-chan amqp091.Delivery {
	newChannel, err := conn.Channel()
	myutils.FailOnError(err, "Failed to open a new channel for consuming messages")

	msgs, err := newChannel.Consume(
		queueName,
		"consumer", // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	myutils.FailOnError(err, "Failed to register a consumer")
	return msgs
}

// DeclareAndBindQueue declares a queue and binds it to a routing key
func DeclareAndBindQueue(exchangeName, queueName, routingKey string) {
	_, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	myutils.FailOnError(err, "Failed to declare the queue")

	err = ch.QueueBind(
		queueName,
		routingKey,
		exchangeName,
		false,
		nil,
	)
	myutils.FailOnError(err, "Failed to bind the queue")
}

// CloseRabbitMQ closes the connection and channel
func CloseRabbitMQ() {
	if ch != nil {
		ch.Close()
	}
	if conn != nil {
		conn.Close()
	}
}
