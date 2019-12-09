package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	s = append(s, 0)  //nil に対しても append できる 
	printSlice(s)  // => len=1 cap=2 [0]

	s = append(s, 1)  //必要に応じて slice は大きくなっていく
	printSlice(s)  // => len=2 cap=2 [0 1]

	s = append(s, 2, 3, 4)  //一度に1つ以上の要素を append することができる
	printSlice(s)  // => len=5 cap=8 [0 1 2 3 4]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}