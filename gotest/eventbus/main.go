package main

import (
	"fmt"
	"github.com/asaskevich/EventBus"
	"time"
)

func calculator(a int, b int) {
	fmt.Printf("%d\n", a+b)
}

func slowCalculator(a, b int) {
	time.Sleep(3 * time.Second)
	fmt.Printf("%d\n", a+b)
}

//func main() {
//	bus := EventBus.New()
//	bus.SubscribeAsync("main:slow_calculator", slowCalculator, true)
//	bus.SubscribeAsync("main:calculator", calculator, true)
//
//	bus.Publish("main:slow_calculator", 20, 40)
//	bus.Publish("main:calculator", 30, 40)
//
//	//bus.WaitAsync() // wait for all async callbacks to complete
//
//	bus.Unsubscribe("main:calculator", calculator)
//	select {}
//}

func main() {
	go func() {
		server := EventBus.NewServer(":2010", "/_server_bus_", EventBus.New())
		server.Start()
		time.Sleep(time.Second * 3)

		// ...
		server.EventBus().Publish("main:calculator", 4, 6)
		// ...
		//server.Stop()
	}()

	go func() {
		client := EventBus.NewClient(":2015", "/_client_bus_", EventBus.New())
		client.Start()
		client.Subscribe("main:calculator", calculator, ":2010", "/_server_bus_")
		// ...
		//client.Stop()
		time.Sleep(time.Second * 10)
	}()
	select {}
}
