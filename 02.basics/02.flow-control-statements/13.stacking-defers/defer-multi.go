package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i ++ {
		defer fmt.Println(i)  //last-in-first-out（ 9-8-7-6-5-4-3-2-1-0 ）で実行
	}

	fmt.Println("done")
}