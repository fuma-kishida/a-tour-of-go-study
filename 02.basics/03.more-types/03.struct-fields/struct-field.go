package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2} //構造体の初期化方法① => {} で順番にフィールドの値を渡す
	v.X = 4  //構造体の初期化方法② => 変数定義後にフィールドを設定する
	fmt.Println(v.X)  //=> 4
}