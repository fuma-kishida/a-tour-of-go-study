package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {  //メソッドが変数レシーバの場合、呼び出し時に、変数 or ポインタいずれかのレシーバとしてとれる => この場合、 p.Abs() は (*p).Abs() として解釈される
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())  // => 5
	fmt.Println(AbsFunc(v))  // => 5

	p := &Vertex{4, 3} 
	fmt.Println(p.Abs())  // => 5
	fmt.Println(AbsFunc(*p))  // => 5
}