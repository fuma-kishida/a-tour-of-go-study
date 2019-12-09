package main

import "fmt"

func main() {
	var i interface{}  //ゼロ個のメソッドを指定されたインターフェース型 => 空のインターフェース => 任意の型の値を保持できる => 全ての型は、少なくともゼロ個のメソッドを実装している => 空のインターフェースは、未知の型の値を扱うコードで使用される
	describe(i)  // => (<nil>, <nil>)

	i = 42
	describe(i)  // => (42, int)

	i = "hello"
	describe(i)  // => (hello, string)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}