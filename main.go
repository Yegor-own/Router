package main

import (
	rabbit "router/rabbit/src"
	switcher "router/switcher/src"
)

func main() {
	var ch chan byte
	go rabbit.Api()
	go switcher.Switcher()
	<-ch
}
