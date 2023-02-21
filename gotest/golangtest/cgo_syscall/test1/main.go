package main

import "C"
import (
	"fmt"
	"sync"
)

//#include <stdio.h>
//int sum(int a, int b) { return a+b; }
//void cprint(char* c) {printf("%s\n", c);}
import "C"

var (
	count = 0
	m     sync.Mutex
)

func main() {
	N := 1000
	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		index := i
		go func() {
			m.Lock()
			count++
			defer wg.Done()
			fmt.Printf("%d\n", C.sum(0, C.int(index)))
			m.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("count = ", count)
}
