package main 

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // => {1, 2}
	v2 = Vertex{X: 1}  // => {1, 0}
	v3 = Vertex{}  // => {0, 0}
	p = &Vertex{1, 2}  // => 新しく割り当てられたstructへのポインタを戻す
)

func main() {
	fmt.Println(v1, p, v2, v3)
}