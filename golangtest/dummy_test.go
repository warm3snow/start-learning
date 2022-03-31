package golangtest

import (
	"fmt"
	"io"
	"math"
	"testing"

	"github.com/pkg/errors"
)

func TestSqrt(t *testing.T) {
	math.Sqrt(10)
	fmt.Println(io.EOF)
	io.EOF = nil
	fmt.Println(io.EOF)
	fmt.Errorf("%w", errors.New("fmtError"))

	//ioutil.ReadAll()
}

func TestMap(t *testing.T) {
	testMap := make(map[string]string)

	fmt.Println(len(testMap["123"]))
}
