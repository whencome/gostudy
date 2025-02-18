// 一般数学函数
// 注意： golang并没有提供求最大公约数的gcd函数，需要自己实现
package main

import (
	"fmt"
	"math"
)

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	fmt.Printf("ceil(2.1) = %f\n", math.Ceil(2.1))
	fmt.Printf("ceil(2.9) = %f\n", math.Ceil(2.9))
	fmt.Printf("ceil(2.1) = %f\n", math.Ceil(-2.1))
	fmt.Printf("ceil(2.9) = %f\n", math.Ceil(-2.9))
	fmt.Printf("floor(2.1) = %f\n", math.Floor(2.1))
	fmt.Printf("floor(2.9) = %f\n", math.Floor(2.9))
	fmt.Printf("floor(2.1) = %f\n", math.Floor(-2.1))
	fmt.Printf("floor(2.9) = %f\n", math.Floor(-2.9))
	fmt.Printf("pow(2,3) = %f\n", math.Pow(2, 3))
	fmt.Printf("pow(2,-3) = %f\n", math.Pow(2, -3))
	fmt.Printf("sqrt(2) = %f\n", math.Sqrt(2))

	fmt.Printf("gcd(128, 456) = %d\n", GCD(128, 456))
}
