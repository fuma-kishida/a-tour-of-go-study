package main 

import (
	"fmt"
	"math"
) 

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {  //さっきのメソッドと同様な処理を関数によって書き直した
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {  //さっきのメソッドと同様な処理を関数によって書き直した
	v.X = v.X * f 
	v.Y = v.Y * f 
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}