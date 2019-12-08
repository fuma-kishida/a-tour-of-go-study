package main 

import "fmt"

//戻り値となる変数に名前をつけることができる
//この場合、return ステートメントに何も書かずに戻せる
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}