package main

import (
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

func main() {
	config := RabbitMQConfig{
		Host:     "localhost",
		Port:     "5672",
		Username: "guest",
		Password: "guest",
	}

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", config.Username, config.Password, config.Host, config.Port))
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Perform RabbitMQ operations here

	fmt.Println("RabbitMQ connection established")

	// Sleep for a while to keep the program running
	time.Sleep(5 * time.Second)
}
