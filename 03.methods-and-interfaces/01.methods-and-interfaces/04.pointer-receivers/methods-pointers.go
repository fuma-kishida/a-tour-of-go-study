package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	//ポインタレシーバにすることで、main関数内で v の値を直接書き換えることができる
	//ポインタレシーバでなく変数レシーバ（=> func (v Vertex) Scale(...)）の場合、ScalaメソッドはVertex変数のコピーを操作するため、main関数内で v の値を書き換えることができない
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}  //Vertex型のstructが定義される
	v.Scale(10)  //(変数レシーバのとき)v のコピーに対して、Scalaメソッドを作用させる //(ポインタレシーバのとき)v に対してScalaメソッドを直接作用させる
	fmt.Println(v.Abs())  // => (変数レシーバのとき) 5 // => (ポインタレシーバのとき) 50
}