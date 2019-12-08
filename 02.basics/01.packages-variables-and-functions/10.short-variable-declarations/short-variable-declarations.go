package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3  //関数内では、var宣言の代わりに、:= による代入文を使って、暗黙的な型宣言ができる
	c, python, java := true, false, "no!"  //関数外では、キーワード(var や func 等)で始まる宣言が必要 => := による暗黙的な宣言ができない

	fmt.Println(i, j, k, c, python, java)
}