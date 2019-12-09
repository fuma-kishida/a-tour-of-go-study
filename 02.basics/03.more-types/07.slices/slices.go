package main

import "fmt"

func main() {
	//array を定義 
	//array は固定長
	primes := [6]int{2, 3, 5, 7, 11, 13}  

	//slice を定義
	//[]T型 は型 T のスライスを表す
	//slice は可変長 
	var s []int = primes[1:4]  
	fmt.Println(s)
}