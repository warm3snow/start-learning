package basic

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"testing"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm3"
)

func Benchmark_SM3_Parallel(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sm3.Sm3Sum(msg)
		}
	})
}

func Benchmark_SHA256_Parallel(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sha256.Sum256(msg)
		}
	})
}

func Benchmark_SM2_Sign_Parallel(b *testing.B) {
	digest := sha256.Sum256(msg)
	priv, _ := sm2.GenerateKey(rand.Reader)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			priv.Sign(rand.Reader, digest[:], nil)
		}
	})
}

func Benchmark_p256_Sign_Parallel(b *testing.B) {
	digest := sha256.Sum256(msg)
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			priv.Sign(rand.Reader, digest[:], nil)
		}
	})
}

func Benchmark_SM2_Verify_Parallel(b *testing.B) {
	digest := sha256.Sum256(msg)
	priv, _ := sm2.GenerateKey(rand.Reader)
	sig, _ := priv.Sign(rand.Reader, digest[:], nil)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			priv.Verify(digest[:], sig)
		}
	})
}

func Benchmark_p256_Verify_Parallel(b *testing.B) {
	digest := sha256.Sum256(msg)
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	sig, _ := priv.Sign(rand.Reader, digest[:], nil)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ecdsa.VerifyASN1(&priv.PublicKey, digest[:], sig)
		}
	})
}
