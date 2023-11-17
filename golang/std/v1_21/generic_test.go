/**
 * @Author: xueyanghan
 * @File: generic_test.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/10/12 17:05
 */

package v1_21

import (
	"fmt"
	"maps"
	"slices"
	"testing"
)

func TestCloneMap(t *testing.T) {
	var intMap = make(map[int]int)
	for i := 0; i < 1000; i++ {
		intMap[i] = i
	}

	intMap2 := maps.Clone(intMap)

	fmt.Println(intMap)
	fmt.Println(intMap2)

	if maps.EqualFunc(intMap, intMap2, func(i, j int) bool {
		return intMap[i] == intMap2[j]
	}) {
		fmt.Println("equal")
	}

}

func TestCloneSlice(t *testing.T) {
	var intSlice []int
	intSlice = append(intSlice, 1)
	intSlice = append(intSlice, 2)

	intSlice2 := slices.Clone(intSlice)

	fmt.Println(intSlice)
	fmt.Println(intSlice2)
}
