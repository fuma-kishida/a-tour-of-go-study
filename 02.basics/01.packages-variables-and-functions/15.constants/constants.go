package main

import "fmt"

const Pi = 3.14  //定数は const キーワードを使って宣言する

func main() {
	const World = "世界"  //定数は、文字、文字列、boolean, 数値のみで使える
	fmt.Println("Hello", World)  //定数は := によって宣言できない
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}