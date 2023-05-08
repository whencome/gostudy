/*
 * 本示例主要演示向量的常见操作.
 */
package main

import (
	"fmt"

	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/mat"
)

func main() {
	// 创建向量
	vectorA := mat.NewVecDense(3, []float64{11.0, 5.2, -1.3})
	vectorB := mat.NewVecDense(3, []float64{-7.2, 4.2, 5.1})
	fmt.Printf("vector A = %+v\n", vectorA)
	fmt.Printf("vector B = %+v\n", vectorB)

	// 1. 向量的加法运算，用AddVec实现，根据定义，为了不覆盖A、B两个向量，我们应当先创建一个空向量
	vectorC := mat.NewVecDense(3, nil)
	vectorC.AddVec(vectorA, vectorB)
	fmt.Printf("vector C = %+v\n", vectorC)

	// 2. 向量的减法运算，用SubVec实现，同加法一样，应当先创建一个空向量
	vectorD := mat.NewVecDense(3, nil)
	vectorD.SubVec(vectorA, vectorB)
	fmt.Printf("vector D = %+v\n", vectorD)

	// 3. 向量的点积
	dotProduct := mat.Dot(vectorA, vectorB)
	fmt.Printf("the dot product of A and B is : %.2f\n", dotProduct)

	// 4. 向量的标量乘法（标量乘以向量）
	vectorA.ScaleVec(1.5, vectorA)
	fmt.Printf("Scaling A by 1.5v gives: %v\n", vectorA)

	// 5. 两个向量的元素积
	vectorE := mat.NewVecDense(3, nil)
	vectorE.MulElemVec(vectorA, vectorB)
	fmt.Printf("vector E = %+v\n", vectorE)

	// 求向量B的欧几里德范数
	normB := blas64.Nrm2(vectorB.RawVector())
	fmt.Printf("The norm/length of B is: %0.2f\n", normB)
}
