package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	// 1. 创建一个3*3的矩阵
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 参数分别是行数、列数以及数据
	matrix := mat.NewDense(3, 3, data)
	// 打印矩阵信息
	fmt.Printf("matrix => %v\n", matrix)

	// 2. 获取第2行第2个值
	// 注意：行和列的下标都是从0开始的
	v22 := matrix.At(1, 1)
	fmt.Printf("v22 => %v\n", v22)

	// 3. 获取第二列的值
	col2 := mat.Col(nil, 1, matrix)
	fmt.Printf("col2 => %v\n", col2)

	// 4. 获取第一行的值
	row1 := mat.Row(nil, 0, matrix)
	fmt.Printf("row1 => %v\n", row1)

	// 5. 修改单个值，修改第3行第1个值
	matrix.Set(2, 0, 7.9)
	fmt.Printf("matrix => %v\n", matrix)

	// 6. 修改一整行的数据，修改第二行数据
	matrix.SetRow(1, []float64{12.65, 3.12, 68.99})
	fmt.Printf("matrix => %v\n", matrix)

	// 7. 修改一整列的数据，修改第3列的数据
	matrix.SetCol(2, []float64{7.88, 9.12, -23.45})
	fmt.Printf("matrix => %v\n", matrix)

}
