package main

import "fmt"

func main() {
	defer fmt.Println("world")  //"hello" が実行されるのを待った後処理される (=> defer)

	fmt.Println("hello")
}