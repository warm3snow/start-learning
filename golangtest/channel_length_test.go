package golangtest

import "testing"

func TestChannel_length(t *testing.T) {
	testChan := make(chan int, 10)
	t.Log("1. chan length = ", len(testChan))

	testChan <- 1
	testChan <- 1
	t.Log("2. chan length = ", len(testChan))

	<-testChan
	t.Log("3. chan length = ", len(testChan))

	close(testChan)
	t.Log("4. chan length = ", len(testChan))

	<-testChan
	t.Log("5. chan length = ", len(testChan))
}
