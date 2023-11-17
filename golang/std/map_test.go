/**
 * @Author: xueyanghan
 * @File: map_test.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/8/23 18:54
 */

package std

import "testing"

type Session struct {
	Name string
}

func TestMapNil(t *testing.T) {
	var m map[string]map[string]*Session
	m = make(map[string]map[string]*Session)

	t.Log("m[a] = ", m["a"])
	t.Log("m[a][b] = ", m["a"]["b"])

	if m["a"] == nil {
		t.Log("m[a] is nil")
	} else {
		t.Log("m[a] is not nil")
	}

	if m["a"]["b"] == nil {
		t.Log("m[a][b] is nil")
	} else {
		t.Log("m[a][b] is not nil")
	}

	if _, ok := m["a"]; ok {
		t.Log("a is ok")
	} else {
		t.Log("a is not ok")
	}

	if _, ok := m["a"]["b"]; ok {
		t.Log("a.b is ok")
	} else {
		t.Log("a.b is not ok")
	}
}
