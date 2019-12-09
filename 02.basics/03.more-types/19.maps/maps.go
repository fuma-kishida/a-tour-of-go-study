package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex  //map を宣言 => string - Vertex(struct) を紐付ける

func main() {
	m = make(map[string]Vertex)  //map を生成 => 指定した型・初期化されたマップが返される
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])  // => {40.68433 -74.39967}
}