package std

import (
	"fmt"
	"testing"
)

var (
	a = b + 1
	b = 1
)

func TestVarInit_outside(t *testing.T) {
	fmt.Println(a)
	fmt.Println(b)
}

//can't be compile
func TestVarInit_inside(t *testing.T) {
	//var (
	//	c int = d + 1
	//	d int = 1
	//)
	//fmt.Println(c)
	//fmt.Println(d)

	//output:
	//./var_init_test.go:21:11: undefined: d
}
