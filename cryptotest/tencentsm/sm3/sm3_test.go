package sm3

import (
	"github.com/stretchr/testify/assert"
	"github.com/tjfoc/gmsm/sm3"
	"testing"
)

var (
	msg = []byte("hello chainmaker")
)

func TestSM3Std(t *testing.T) {
	h1 := New()
	_, err := h1.Write(msg)
	assert.NoError(t, err)
	digest1 := h1.Sum(nil)

	h2 := sm3.New()
	_, err = h2.Write(msg)
	assert.NoError(t, err)
	digest2 := h2.Sum(nil)

	assert.Equal(t, digest1, digest2)
}
