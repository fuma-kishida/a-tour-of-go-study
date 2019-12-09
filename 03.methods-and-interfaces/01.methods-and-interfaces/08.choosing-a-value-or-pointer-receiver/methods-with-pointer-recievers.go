package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/*
ポインタレシーバを使う目的は2つ
	- メソッドがレシーバが指す先の変数を変更するため
	- メソッドの呼び出し毎に変数のコピーを避けるため
*/
func (v *Vertex) Scale(f float64) {  //ポインタレシーバ
	v.X = v.X * f
	v.Y = v.Y * f  
}

func (v *Vertex) Abs() float64 {  //ポインタレシーバ
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}  //アドレス
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}