package main

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

func main() {
	// 创建2个3*3的矩阵
	a := mat.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})
	b := mat.NewDense(3, 3, []float64{8, 9, 10, 1, 4, 2, 9, 0, 2})

	// 创建一个不同规格的矩阵
	c := mat.NewDense(3, 2, []float64{3, 2, 1, 4, 0, 8})

	// 计算矩阵的加法，a+b
	var d mat.Dense
	d.Add(a, b)
	fd := mat.Formatted(&d)
	fmt.Printf("d = a + b = \n%.04v\n", fd)

	// 矩阵相乘
	var f mat.Dense
	f.Mul(a, c)
	ff := mat.Formatted(&f)
	fmt.Printf("f = a * c = \n%0.4v\n", ff)

	// raising a matrix to a power
	var g mat.Dense
	g.Pow(a, 5)
	fg := mat.Formatted(&g)
	fmt.Printf("g = a ^ 5 = \n%0.4v\n", fg)

	// 对每个元素应用一个方法
	var h mat.Dense
	sqrt := func(_, _ int, v float64) float64 {
		return math.Sqrt(math.Abs(v))
	}
	h.Apply(sqrt, a)
	fh := mat.Formatted(&h)
	fmt.Printf("h = sqrt(a) = \n%0.4v\n", fh)
}
