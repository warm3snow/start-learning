package cryptotest

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"testing"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm3"
)

func Benchmark_SM3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm3.Sm3Sum(msg)
	}
}

func Benchmark_SHA256(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sha256.Sum256(msg)
	}
}

func Benchmark_SM2_Sign(b *testing.B) {
	digest := sha256.Sum256(msg)
	priv, _ := sm2.GenerateKey(rand.Reader)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		priv.Sign(rand.Reader, digest[:], nil)
	}
}

func Benchmark_p256_Sign(b *testing.B) {
	digest := sha256.Sum256(msg)
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		priv.Sign(rand.Reader, digest[:], nil)
	}
}

func Benchmark_SM2_Verify(b *testing.B) {
	digest := sha256.Sum256(msg)
	priv, _ := sm2.GenerateKey(rand.Reader)
	sig, _ := priv.Sign(rand.Reader, digest[:], nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		priv.Verify(digest[:], sig)
	}
}

func Benchmark_p256_Verify(b *testing.B) {
	digest := sha256.Sum256(msg)
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	sig, _ := priv.Sign(rand.Reader, digest[:], nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ecdsa.VerifyASN1(&priv.PublicKey, digest[:], sig)
	}
}
