package main

import (
	"switcher/rabbit"
)

func main() {
	go rabbit.ConnectRabbitMQ()
}
