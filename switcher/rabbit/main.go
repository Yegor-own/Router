package main

import (
	"io"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	r := gin.Default()
	r.GET("/rabbit/send", func(c *gin.Context) {

		read := c.Request.Body
		msg, err := io.ReadAll(read)
		FailOnError(err, "Failed to read Request Body")
		SendRabbitMQ(ch, []byte(msg))
	})
	r.Run(":2100") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
