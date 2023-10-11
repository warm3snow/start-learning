/**
 * @Author: xueyanghan
 * @File: time_test.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/9/6 11:54
 */

package std

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000000"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000000000"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000000000 -0700 MST"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000000000 -0700 MST Mon"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000000000 -0700 MST Mon Jan"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000000000 -0700 MST Mon Jan 2"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000000000 -0700 MST Mon Jan 2 15"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000000000 -0700 MST Mon Jan 2 15:04"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000000000 -0700 MST Mon Jan 2 15:04:05"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000000000 -0700 MST Mon Jan 2 15:04:05.000"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000000000 -0700 MST Mon Jan 2 15:04:05.000000"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000000000 -0700 MST Mon Jan 2 15:04:05.000000000"))
}

func TestTimeParseAndFormat(t *testing.T) {
	timeStr := "2023-09-05T19:06:53.302+08:00"
	now, err := time.Parse(time.RFC3339, timeStr)
	assert.NoError(t, err)

	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000000"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000000000"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000000000 -0700 MST"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000000000 -0700 MST Mon"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000000000 -0700 MST Mon Jan"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000000000 -0700 MST Mon Jan 2"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000000000 -0700 MST Mon Jan 2 15"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000000000 -0700 MST Mon Jan 2 15:04"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000000000 -0700 MST Mon Jan 2 15:04:05"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000000000 -0700 MST Mon Jan 2 15:04:05.000"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000000000 -0700 MST Mon Jan 2 15:04:05.000000"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000000000 -0700 MST Mon Jan 2 15:04:05.000000000"))
}
