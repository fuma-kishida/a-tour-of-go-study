package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)  // t := i.(T) => インターフェースの値 i が具体的な型 T を保持し、基になる T の値を変数 t に代入することを主張する
	fmt.Println(s)  // => hello

	s, ok := i.(string)
	fmt.Println(s, ok)  // => hello true

	f, ok := i.(float64)
	fmt.Println(f, ok)  // => 0 false

	f = i.(float64)
	fmt.Println(f)  // i が T を保持していない場合 => panic
}