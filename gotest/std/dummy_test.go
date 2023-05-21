package std

import (
	"encoding/hex"
	"fmt"
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

func TestNil(t *testing.T) {
	var err error
	var err2 error
	fmt.Println(err == nil)
	fmt.Println(&err)
	fmt.Println(&err2)

}
