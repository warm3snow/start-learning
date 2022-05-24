package golangtest

import (
	"math"
	"testing"
)

func BenchmarkPow_0_5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Pow(10, 0.5)
	}
}

func BenchmarkSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Sqrt(10)
	}
}
