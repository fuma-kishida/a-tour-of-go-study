package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {  //; までも省略可
		sum += sum
	}
	fmt.Println(sum)
}