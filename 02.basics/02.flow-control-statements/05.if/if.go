package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {  //if 条件式 {処理内容}
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
 
func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}

