package main

import (
	"fmt"
	"math"
)

//interface は関数（メソッド）をまとめたもの
//具体的な実装を気にすることなく1つのまとまった塊として定義できることで、他の関数から扱いやすくなる
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)  // => -1.4142135623730951
	v := Vertex{3, 4}  // => {3, 4}

	a = f   // => -1.4142135623730951
	a = &v  // => &{3 4}

//	a = v  // => error : Vertex does not implement Abser (Abs method has pointer receiver)

	fmt.Println(a.Abs())  // => 5
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
} 

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}