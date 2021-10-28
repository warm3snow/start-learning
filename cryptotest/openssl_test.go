package cryptotest

import (
	"crypto/rand"
	"testing"

	"github.com/tjfoc/gmsm/sm2"
)

func BenchmarkOpensslSM2_Sign(b *testing.B) {
	priv, _ := sm2.GenerateKey(rand.Reader)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		priv.Sign(rand.Reader, msg, nil)
	}
}

func BenchmarkOpensslSM2_Verify(b *testing.B) {
	priv, _ := sm2.GenerateKey(rand.Reader)
	sig, _ := priv.Sign(rand.Reader, msg, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		priv.PublicKey.Verify(msg, sig)
	}
}
