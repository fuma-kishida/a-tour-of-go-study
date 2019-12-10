package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//Go には class という概念は無いが、構造体内に method を定義できる
//Vertex という構造体に対して、Abs というメソッドを定義する
func (v Vertex) Abs() float64 {  //func レシーバ メソッド名 返り値の型 {メソッドが行う処理内容}
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())  //メソッド呼び出し => 5
}