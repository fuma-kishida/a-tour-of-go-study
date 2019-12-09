package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)  // => len=6 cap=6 [2 3 5 7 11 13]

	s = s[:0]
	printSlice(s)  // => len=0 cap=6 []

	s = s[:4]
	printSlice(s)  // => len=4 cap=6 [2 3 5 7]

	s = s[2:]
	printSlice(s)  // => len=2 cap=4 [5 7]
}

func printSlice(s []int) {
	//スライスの長さ => それに含まれる要素の数 => len() で取得できる
	//スライスの容量 => スライスの最初の要素から数えて、もととなる配列の要素数 => cap() で取得できる
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}