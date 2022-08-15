package rabbit

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Send(ch *amqp.Channel, msg []byte, route string) {

	err := ch.Publish(
		"",    // exchange
		route, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(msg),
		})
	FailOnError(err, "Failed to publish a message")
}
