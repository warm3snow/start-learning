package golangtest

import (
	"fmt"
	"testing"
	"time"
)

func TestGoto(t *testing.T) {
	intChan := make(chan int, 1)
	intChan <- 1
	for {
	XLoop:
		for {
			select {
			case <-intChan:
				fmt.Println("5s ticker")
			default:
				break XLoop
			}
			fmt.Println("Xloop end")
			time.Sleep(5 * time.Second)
		}
		fmt.Println("outer for")
		time.Sleep(5 * time.Second)
	}
}
