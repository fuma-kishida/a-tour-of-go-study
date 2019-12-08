package main

import (
	"fmt"
	"math"
)

//変数 v, 型 T => T(v) で 変数 v を T 型に変換できる
func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}