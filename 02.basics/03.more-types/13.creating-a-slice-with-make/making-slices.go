package main

import "fmt"

func main() {
	a := make([]int, 5)  //組み込みの make 関数を使用して、スライス（動的サイズの配列）を作成できる
	printSlice("a", a)  //a len=5 cap=5 [0 0 0 0 0]

	b := make([]int, 0, 5)  //make 関数の 3番目の引数に、スライスの容量を指定できる
	printSlice("b", b)  //b len=0 cap=5 []

	c := b[:2]
	printSlice("c", c)  //c len=2 cap=5 [0 0]

	d := c[2:5]
	printSlice("d", d)  //d len=3 cap=3 [0 0 0]
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
	s, len(x), cap(x), x)
}