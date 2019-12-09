package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)  // v => Vertex{6, 8}
	ScaleFunc(&v, 10)  // v => Vertex{60, 80}

	p := &Vertex{4, 3}  // p => Vertex{4, 3} の アドレス
	p.Scale(3)  // => &Vertex{12, 9}
	ScaleFunc(p, 8)  // => &Vertex{96, 72}

	fmt.Println(v, p)  // => {60 80} &{96 72}
}