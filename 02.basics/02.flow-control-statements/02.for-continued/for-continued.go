package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 1000; {  //初期化 と 後処理ステートメント の記述は任意
		sum += sum
	}
	fmt.Println(sum)
}