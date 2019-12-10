package main

import "fmt"

//struct (構造体) は、オブジェクト指向言語における class と似たようなもので、関連する変数をひとまとめにする
//struct (構造体) は、下記のように type と struct を使って定義する
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}