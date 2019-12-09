package main

import (
	"fmt"
	"math"
)

type Vertex struct {  // => （今更だけど） struct型 に Verterx型 という別名をつけている
	X, Y float64
}

func Abs(v Vertex) float64 {  //さっきのメソッドと同様の機能を通常の関数によって記述している
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}