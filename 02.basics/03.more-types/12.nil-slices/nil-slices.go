package main

import "fmt"

func main() {
	var s []int 
	fmt.Println(s, len(s), cap(s))  //nilスライスは、長さ 0 , 容量 0 のスライス
	if s == nil {
		fmt.Println("nil!")
	}
}