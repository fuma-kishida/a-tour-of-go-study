package main

import "fmt"

// var宣言では変数ごとに初期値を与えられる
var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}