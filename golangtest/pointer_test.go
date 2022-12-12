package golangtest

import (
	"fmt"
	"runtime"
	"testing"
)

type GF struct {
	Age    *int
	Hight  *int
	Weight *int
}

func TestGF(t *testing.T) {
	val := 30

	gfMap := make(map[string]*GF)
	gf1 := &GF{
		Age:    &val,
		Hight:  &val,
		Weight: &val,
	}
	gfMap["gf1"] = gf1

	gf2 := gfMap["gf1"]
	fmt.Printf("%#v\n", gf2)

	gfMap = nil
	fmt.Printf("%#v\n", gf2)

	runtime.GC()
	fmt.Printf("%#v\n", gf2)
}
