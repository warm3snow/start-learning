package main

import "chainmaker.org/gotest/golangtest/dummy"

func hello(msg string) {
	println(msg)
}

func main() {
	go func() {
		for {
			hello("xxx")
		}
	}()
	//time.Sleep(1)
	hello("hello world")

	dummy.DummyPrint()
}
