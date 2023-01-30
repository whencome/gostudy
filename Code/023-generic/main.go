package main

import (
	"fmt"
)

// Number 声明约束
type Number interface {
	int16 | int32 | int64 | int | uint | uint16 | uint32 | uint64 | float32 | float64
}

func sum[T Number](numbers ...T) T {
	var s T
	for _, n := range numbers {
		s += n
	}
	return s
}

func contains[T comparable](dst []T, v T) bool {
	if len(dst) == 0 {
		return false
	}
	for _, t := range dst {
		if t == v {
			return true
		}
	}
	return false
}

func main() {
	g1 := []int{1, 4, 67, 89, 22, 54}
	g2 := []float64{13.45, 67.22, 56.88, 92.41}

	s1 := sum(g1...)
	s2 := sum(g2...)

	fmt.Printf("sum(g1) = %d\n", s1)
	fmt.Printf("sum(g2) = %.2f\n", s2)

	fmt.Printf("g1 contains 89 ? %v\n", contains(g1, 89))
	fmt.Printf("g1 contains 102 ? %v\n", contains(g1, 102))
	fmt.Printf("g2 contains 67.22 ? %v\n", contains(g2, 67.22))
	fmt.Printf("g2 contains 102 ? %v\n", contains(g2, 102))
}
