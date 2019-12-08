package main

import "fmt"

// 関数の2つ以上の引数の型が同じ場合は、最後の方を省略できる
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}