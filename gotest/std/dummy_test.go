package std

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"math"
	"strconv"
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

func TestStrconv(t *testing.T) {
	//var appId int64 = 1000000000
	//appIdStr := strconv.FormatInt(appId, 10)
	//fmt.Println("appIdStr", appIdStr)

	appIdBytes := []byte("1000000000")

	appId, _ := strconv.ParseInt(string(appIdBytes), 10, 64)
	fmt.Println("appId", appId)
}

type SliceTest struct {
	IntList []int `json:"IntList"`
}

func TestSlicePrintf(t *testing.T) {
	var slice []int
	sliceTest1 := &SliceTest{
		IntList: slice,
	}
	sliceTest1JsonBytes, err := json.Marshal(sliceTest1)
	assert.NoError(t, err)
	fmt.Println(string(sliceTest1JsonBytes))

	slice2 := make([]int, 0)
	sliceTest2 := &SliceTest{
		IntList: slice2,
	}
	sliceTest2JsonBytes, err := json.Marshal(sliceTest2)
	assert.NoError(t, err)
	fmt.Println(string(sliceTest2JsonBytes))
}

func TestIntMap(t *testing.T) {
	var intMap map[int]int

	v := intMap[100]
	fmt.Println(v)
}
