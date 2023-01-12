package golangtest

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/davecgh/go-spew/spew"
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

type TestStruct struct {
	IntSlice []*int
}

func TestSlicePointer(t *testing.T) {
	s := make([]*int, 0, 10)
	i1 := 1
	i2 := 2
	i3 := 3
	s = append(s, &i1)
	s = append(s, &i2)
	s = append(s, &i3)
	testStruct := &TestStruct{
		IntSlice: s,
	}

	testSlice := testStruct.IntSlice

	ii := 4
	testSlice[2] = &ii

	spew.Dump(testStruct)
}

func TestSlice(t *testing.T) {
	intSlice := make([]int, 0, 10)
	intSlice = append(intSlice, 1)

	i := 0

	xx := append(intSlice[:i], intSlice[i+1:]...)

	fmt.Println(intSlice[1:])
	fmt.Println(xx)

}
