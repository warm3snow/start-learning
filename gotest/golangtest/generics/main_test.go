/**
	Benchmark conclusion:
	10% performance downside, but not the absolute. sometime we found no effect.
	Why? generics test case has cache? TODO
**/
package main

import (
	"fmt"
	"testing"
)

var ints map[string]int64

func init() {
	ints = make(map[string]int64)
	for i := 0; i < 10000; i++ {
		ints[fmt.Sprintf("%d", i)] = int64(i)
	}
}

func BenchmarkSumInts(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SumInts(ints)
	}
}

func BenchmarkSumIntsOrFloats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumIntsOrFloats(ints)
	}
}

func BenchmarkSumNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumNumbers(ints)
	}
}
