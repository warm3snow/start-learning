package sm2

import (
	"testing"

	"chainmaker.org/gotest/opencrypto/tencentsm/sm3"

	"github.com/spf13/viper"
	tjsm3 "github.com/tjfoc/gmsm/sm3"
	tjx509 "github.com/tjfoc/gmsm/x509"

	"github.com/stretchr/testify/assert"
)

var (
	msg = []byte("hello world")
)

func BenchmarkSignAndVerifyParallel(b *testing.B) {
	h := sm3.New()
	_, err := h.Write(msg)
	assert.NoError(b, err)
	digest := h.Sum(nil)

	priv, pub, err := GenerateKeyPair()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sig, err := priv.Sign(nil, digest, nil)
			assert.NoError(b, err)

			ok := pub.Verify(digest, sig)
			assert.True(b, ok)
		}
	})
}

func TestSignAndVerify(t *testing.T) {
	h := sm3.New()
	_, err := h.Write(msg)
	assert.NoError(t, err)

	digest := h.Sum(nil)

	priv, pub, err := GenerateKeyPair()

	sig, err := priv.Sign(nil, digest, nil)
	assert.NoError(t, err)

	pass := pub.Verify(digest, sig)
	assert.True(t, pass)
}

func BenchmarkSM2_Sign(b *testing.B) {
	h := sm3.New()
	_, err := h.Write(msg)
	assert.NoError(b, err)
	digest := h.Sum(nil)
	priv, _, err := GenerateKeyPair()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		priv.Sign(nil, digest, nil)
	}
}

func BenchmarkSM2_Verify(b *testing.B) {
	h := sm3.New()
	_, err := h.Write(msg)
	assert.NoError(b, err)
	digest := h.Sum(nil)

	viper.Set("common.tencentsm.ctx_pool_size.max", 100)
	viper.Set("common.tencentsm.ctx_pool_size.init", 100)

	priv, pub, err := GenerateKeyPair()
	sig, _ := priv.Sign(nil, digest, nil)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ok := pub.Verify(digest, sig)
		assert.True(b, ok)
	}
}

func BenchmarkSM2_Verify_Parallel(b *testing.B) {
	h := sm3.New()
	_, err := h.Write(msg)
	assert.NoError(b, err)
	digest := h.Sum(nil)
	priv, pub, err := GenerateKeyPair()
	sig, _ := priv.Sign(nil, digest, nil)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ok := pub.Verify(digest, sig)
			assert.True(b, ok)
		}
	})
}

func BenchmarkTencentSM3(b *testing.B) {
	h := sm3.New()
	_, err := h.Write(msg)
	assert.NoError(b, err)
	for i := 0; i < b.N; i++ {
		h.Sum(nil)
	}
}

func TestSM2Std(t *testing.T) {
	h := sm3.New()
	_, err := h.Write(msg)
	assert.NoError(t, err)
	digest := h.Sum(nil)

	for i := 0; i < 300; i++ {
		priv, _, err := GenerateKeyPair()
		sig, err := priv.Sign(nil, digest, nil)
		assert.NoError(t, err)

		sm2Priv, _ := MarshalPrivateKey(priv)
		privateKey, _ := tjx509.ParsePKCS8UnecryptedPrivateKey(sm2Priv)

		pass := privateKey.Verify(digest, sig)
		assert.True(t, pass)
	}
}

func TestSM3Std(t *testing.T) {
	h := sm3.New()
	_, err := h.Write(msg)
	assert.NoError(t, err)
	digest1 := h.Sum(nil)

	th := tjsm3.New()
	_, err = th.Write(msg)
	assert.NoError(t, err)
	digest2 := th.Sum(nil)

	assert.Equal(t, digest2, digest1)
}
