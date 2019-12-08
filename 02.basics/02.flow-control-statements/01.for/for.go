package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {  //for 初期化ステートメント; 条件式; 後処理ステートメント {処理内容}
		sum += i
	}
	fmt.Println(sum)
}