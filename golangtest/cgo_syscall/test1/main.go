package main

import "C"
import (
	"fmt"
	"runtime"
	"sync"
)

//#include <stdio.h>
//int sum(int a, int b) { return a+b; }
//void cprint(char* c) {printf("%s\n", c);}
import "C"

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			fmt.Printf("%d\n", C.sum(0, C.int(i)))
			wg.Done()
		}(i, &wg)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			C.cprint(C.CString("hello world"))
			wg.Done()
			runtime.Gosched()
		}(i, &wg)
	}
	wg.Wait()
}
