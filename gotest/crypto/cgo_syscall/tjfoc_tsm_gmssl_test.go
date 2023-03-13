package cryptotest

import (
	"fmt"
	"testing"

	gsm2 "chainmaker.org/chainmaker/common/v2/opencrypto/gmssl/sm2"
	tsm2 "chainmaker.org/chainmaker/common/v2/opencrypto/tencentsm/sm2"

	"github.com/stretchr/testify/assert"
)

var msg = []byte("hello-world")

func testGsm2(t *testing.T) {
	priv, err := gsm2.GenerateKeyPair()
	assert.NoError(t, err)
	sig, err := priv.Sign(msg)
	assert.NoError(t, err)
	for i := 0; i < 100; i++ {
		ok, err := priv.PublicKey().Verify(msg, sig)
		assert.NoError(t, err)
		assert.True(t, ok)
	}
}
func testTsm2(t *testing.T) {
	priv, err := tsm2.GenerateKeyPair()
	assert.NoError(t, err)
	sig, err := priv.Sign(msg)
	assert.NoError(t, err)

	for i := 0; i < 100; i++ {
		ok, err := priv.PublicKey().Verify(msg, sig)
		assert.NoError(t, err)
		assert.True(t, ok)
	}
}

func TestCgoParallel(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("GSM2-%d", i+1), testGsm2)
	}
	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("TSM2-%d", i+1), testTsm2)
	}
}

func TestCgoParallel2(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("GSM2-%d", i+1), testGsm2)
	}
	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("TSM2-%d", i+1), testTsm2)
	}
}
