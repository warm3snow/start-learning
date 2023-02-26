package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	ints := map[string]int64{
		"first":  11,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  11.00,
		"second": 12.00,
	}

	fmt.Printf("Non-generics Sums: %v and %v\n", SumInts(ints), SumFloat64(floats))

	fmt.Printf("Generics Sums: %v and %v\n", SumIntsOrFloats[string, int64](ints), SumIntsOrFloats[string, float64](floats))
	fmt.Printf("Generics Sums(omit types): %v and %v\n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))

	fmt.Printf("Generics Sums(omit types, type constraints): %v and %v\n", SumNumbers(ints), SumNumbers(floats))
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloat64(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// func SumIntsOrFloat64[K comparable, V int64 | float64](m map[K]V) V {
// 	var s V
// 	for _, v := range m {
// 		s += v
// 	}
// 	return s
// }
// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
