package std

import (
	"fmt"
	"testing"
)

func TestPanics(t *testing.T) {
	A()
}

func A() {
	fmt.Println("A start")
	defer func() {
		x := recover() //can recover C panic, when panic, only execute this defer-recover func
		fmt.Println(x.(struct{ ErrStr string }).ErrStr)
	}()

	B()

	fmt.Println("A end")
}

func B() {
	fmt.Println("B start")
	C()
	fmt.Println("B enc")
}

func C() {
	fmt.Println("C start")

	panic(struct {
		ErrStr string
	}{"C"})

	fmt.Println("C end")
}
