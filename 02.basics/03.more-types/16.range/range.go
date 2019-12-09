package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	//for 文で使われる range はスライスやマップをひとつずつ反復処理するために使う
	for i, v := range pow {  //スライスを range で繰り返す場合、range 反復ごとに 2つの変数を返す => 1つ目の変数はインデックスで 2つ目の変数はインデックスの場所の要素のコピー  
		fmt.Printf("2**%d = %d\n", i, v)  
	}
}