package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},  //mapに渡すトップレベルの型が単純な型名である場合、型名を省略できる
	"Google": {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}