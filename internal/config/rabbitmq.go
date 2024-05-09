package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func ConnectRabbitMQ() (*amqp.Connection, error) {

	_ = godotenv.Load()

	url := os.Getenv("RABBITMQ_URL")

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		return nil, err
	}
	return conn, nil
}
