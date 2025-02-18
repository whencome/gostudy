// 三角函数
// 注意：golang并没有提供角度和弧度相互转换的函数，需要自己实现
package main

import (
	"fmt"
	"math"
)

// 将角度转换为弧度
// 公式： 弧度=角度× π / 180
func DegreesToRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}

// 将弧度转换为角度
func RadiansToDegrees(radians float64) float64 {
	return radians * (180 / math.Pi)
}

func main() {
	fmt.Printf("DegreesToRadians(30) = %v\n", DegreesToRadians(30))
	fmt.Printf("sin(DegreesToRadians(30)) = %v\n", math.Sin(DegreesToRadians(30)))
}
