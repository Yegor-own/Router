package rabbit

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"io"
	"net/http"
)

func Api() {
	fmt.Println("Rabbit API running")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	http.HandleFunc("/rabbit/send", func(writer http.ResponseWriter, request *http.Request) {
		read := request.Body
		msg, err := io.ReadAll(read)
		FailOnError(err, "Failed to read Request Body")
		Send(ch, []byte(msg), "switcher")
		writer.Write([]byte("Message added to queue"))
		fmt.Println("Message added to queue")
	})
	http.ListenAndServe(":2100", nil)
	//r := gin.Default()
	//r.GET("/rabbit/send", func(c *gin.Context) {
	//
	//	read := c.Request.Body
	//	msg, err := io.ReadAll(read)
	//	FailOnError(err, "Failed to read Request Body")
	//	Send(ch, []byte(msg), "switcher")
	//})
	//r.Run(":2100") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
