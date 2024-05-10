package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func ConnectMQ() (*amqp.Connection, *amqp.Channel) {
	_ = godotenv.Load()
	url := os.Getenv("RABBITMQ_URL")

	conn, err := amqp.Dial(url)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()
	
	return conn, ch
}
