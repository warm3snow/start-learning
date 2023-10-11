/**
 * @Author: xueyanghan
 * @File: nil_test.og.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/8/30 16:03
 */

package std

import (
	"fmt"
	"testing"
)

func TestNilAssertType(t *testing.T) {
	var a interface{}

	switch a.(type) {
	case nil:
		println("a is nil")
	default:
		println("a is not nil")
	}
}

func TestNil(t *testing.T) {
	var err error
	var err2 error
	fmt.Println(err == nil)
	fmt.Println(&err)
	fmt.Println(&err2)

}
