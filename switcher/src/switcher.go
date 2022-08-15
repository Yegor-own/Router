package switcher

import "fmt"

func Switcher() {
	fmt.Println("Switcher running")
	res := make(chan Message)
	defer close(res)
	go func() {
		for msg := range res {
			fmt.Println(msg)
		}
	}()
	ReceiveRabbitMQ(res)
}
