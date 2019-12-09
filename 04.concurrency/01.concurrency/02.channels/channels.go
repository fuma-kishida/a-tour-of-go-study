package main 

import "fmt"

//チャネル( Channel )型は、チャネルオペレータの <- を用いて値の送受信ができる通り道
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum  // sum を チャネル c へ送信する
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)  // チャネルを生成
	go sum(s[:len(s)/2], c)  //スライス内の数値を合算し、2つの goroutine 間で作業を分配 => 両方の goroutine で計算が完了すると、最終結果が計算される
	go sum(s[len(s)/2:], c)  //スライス内の数値を合算し、2つの goroutine 間で作業を分配 => 両方の goroutine で計算が完了すると、最終結果が計算される
	x, y := <-c, <-c  // c から受信した変数を x, y へ割り当てる

	fmt.Println(x, y, x+y)
}
