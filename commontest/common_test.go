package commontest

import (
	"fmt"
	"runtime"
	"testing"
)

func TestOS(t *testing.T) {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
