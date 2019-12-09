package main 

import (
	"fmt"
	"math"
)

//関数も変数 => 他の変数のように、関数を渡すことができる
//関数値（function-value）は関数の引数に取ることもできるし、戻り値としても利用できる
func compute(fn func(float64, float64) float64) float64 {  
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}