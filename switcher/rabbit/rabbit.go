package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func SendRabbitMQ(ch *amqp.Channel, msg []byte) {

	err := ch.Publish(
		"",       // exchange
		"Router", // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(msg),
		})
	FailOnError(err, "Failed to publish a message")
}
