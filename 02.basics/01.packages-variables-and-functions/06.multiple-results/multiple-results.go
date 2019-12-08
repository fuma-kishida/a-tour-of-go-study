package main

import "fmt"

//関数は複数の戻り値を返すことができる
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}