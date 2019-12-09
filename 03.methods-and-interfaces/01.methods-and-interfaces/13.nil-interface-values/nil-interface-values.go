package main 

import "fmt"

type I interface {
	M()
}

func main() {
	var i I  //interface型の変数を宣言
	describe(i)  // => (<nil>, <nil>)
	i.M()  // => nilインターフェースのメソッドを呼び出すと、ランタイムエラー
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}