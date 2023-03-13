package std

import (
	"encoding/hex"
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

func TestMakeMap(t *testing.T) {
	map1 := make(map[string]string, 2)
	map2 := make(map[string]string)

	fmt.Println(len(map1))
	fmt.Println(len(map2))
}

func TestHexDump(t *testing.T) {
	fmt.Println(hex.Dump([]byte("hello world")))
}
