package main

import (
	"fmt"

	"gonum.org/v1/gonum/floats" // 本示例的向量操作使用floats实现
)

func main() {
	// 使用浮点数初始化两个向量
	vectorA := []float64{11.0, 5.2, -1.3}
	vectorB := []float64{-7.2, 4.2, 5.1}

	// 计算两个向量的点积
	dotProduct := floats.Dot(vectorA, vectorB)
	fmt.Printf("the dot product of A and B is : %.2f\n", dotProduct)

	// 向量的标量乘法（标量乘以向量）
	floats.Scale(1.5, vectorA)
	fmt.Printf("Scaling A by 1.5v gives: %v\n", vectorA)

	// 求向量B的欧几里德范数
	normB := floats.Norm(vectorB, 2)
	fmt.Printf("The norm/length of B is: %0.2f\n", normB)
}
