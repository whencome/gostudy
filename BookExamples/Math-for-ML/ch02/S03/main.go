// log()函数
// golang没有提供任意底的对数的方法，需要使用换底公式自行实现
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("log(4) = %.6f\n", math.Log(4))
	fmt.Printf("log(100) = %.6f\n", math.Log(100))
	fmt.Printf("log2(4) = %.6f\n", math.Log2(4))
	fmt.Printf("log10(100) = %.6f\n", math.Log10(100))
	fmt.Printf("log(e) = %.6f\n", math.Log(math.E))
	// 求以2为底4的对数
	fmt.Printf("log(2, 4) = %.6f\n", math.Log(4)/math.Log(2))
	// 求以3为底27的对数
	fmt.Printf("log(3, 27) = %.6f\n", math.Log(27)/math.Log(3))
}
