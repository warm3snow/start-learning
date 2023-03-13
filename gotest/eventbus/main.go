package main

import (
	"fmt"
	"github.com/asaskevich/EventBus"
)

func cal(a, b int) {
	fmt.Printf("%d\n", a+b)
}

func main() {
	bus := EventBus.New()
	bus.SubscribeAsync("slow::cal", cal, false)

	bus.Publish("slow::cal", 1, 2)
}

func Server() {

}