package golangtest

import (
	"fmt"
	"reflect"
	"testing"
)

type People struct {
	Age  int
	Nmae string
}

func TestReflectStruct(t *testing.T) {
	hxy := &People{18, "hxy"}

	val := reflect.TypeOf(*hxy)
	for i := 0; i < val.NumField(); i++ {
		fmt.Printf("%s: %s\n", val.Field(i).Name, val.Field(i).Type)
	}
}
