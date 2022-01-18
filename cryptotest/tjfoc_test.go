package cryptotest

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	gsm2 "chainmaker.org/chainmaker/common/v2/opencrypto/gmssl/sm2"
	"github.com/tjfoc/gmsm/sm2"
)

func BenchmarkTjfocSM2_Sign(b *testing.B) {
	priv, _ := sm2.GenerateKey(rand.Reader)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		priv.Sign(rand.Reader, msg, nil)
	}
}

func BenchmarkTjfocSM2_Verify(b *testing.B) {
	priv, _ := sm2.GenerateKey(rand.Reader)
	sig, _ := priv.Sign(rand.Reader, msg, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		priv.PublicKey.Verify(msg, sig)
	}
}

func BenchmarkEcdsaP256_Verify(b *testing.B) {
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	sig, _ := priv.Sign(rand.Reader, msg, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ok := ecdsa.VerifyASN1(&priv.PublicKey, msg, sig)
		assert.True(b, ok)
	}
}

func BenchmarkGmsslSM2_verify(b *testing.B) {
	priv, _ := gsm2.GenerateKeyPair()
	sig, _ := priv.Sign(msg)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ok, err := priv.PublicKey().Verify(msg, sig)
		assert.NoError(b, err)
		assert.True(b, ok)
	}
}
