package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {  //インデックスだけが必要な場合は、 , value を省略する
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {  //インデックスや値は、 _ へ代入することで捨てられる
		fmt.Printf("%d\n", value)
	}
}