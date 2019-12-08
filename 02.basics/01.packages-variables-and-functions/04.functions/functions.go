package main

import "fmt"

// 関数は 0個以上の引数を取れる
// 引数の後ろに型をかく
func add(x int, y int) int{
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}