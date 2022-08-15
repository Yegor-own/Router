package switcher

import (
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func ReceiveRabbitMQ(res chan Message) {
	fmt.Println("ReceiveRabbitMQ running")
	defer fmt.Println("ReceiveRabbitMQ stopped")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"switcher", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	FailOnError(err, "Failed to register a consumer")

	var forever2 chan bool
	defer close(forever2)

	go func() {
		var msg Message
		for d := range msgs {
			err := json.Unmarshal(d.Body, &msg)
			FailOnError(err, "Failed to unmarshal json")
			res <- msg
		}
	}()

	<-forever2
}
