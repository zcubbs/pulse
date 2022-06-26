package utils

import (
	"log"
	"os"
)
import "github.com/streadway/amqp"

var channelRabbitMQ *amqp.Channel

func ConnectRabbitmq() (*amqp.Connection, *amqp.Channel) {
	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(getAmqpServerUrl())
	if err != nil {
		panic(err)
	}

	// Let's start by opening a channel to our RabbitMQ
	// instance over the connection we have already
	// established.
	channelRabbitMQ, err = connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}

	// With the instance and declare Queues that we can
	// publish and subscribe to.
	_, err = channelRabbitMQ.QueueDeclare(
		"pulse_pipeline_status_events.input", // queue name
		true,                                 // durable
		false,                                // auto delete
		false,                                // exclusive
		false,                                // no wait
		nil,                                  // arguments
	)
	if err != nil {
		panic(err)
	}
	log.Println("Successfully connected to RabbitMQ")

	return connectRabbitMQ, channelRabbitMQ
}

func GetChannelRabbitMQ() *amqp.Channel {
	return channelRabbitMQ
}

func getAmqpServerUrl() string {
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")
	if amqpServerURL == "" {
		amqpServerURL = "amqp://localhost:5672"
	}

	return amqpServerURL
}
