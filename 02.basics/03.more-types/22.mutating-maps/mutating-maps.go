package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42  //m[key] = elem => m への要素の挿入
	fmt.Println("The value:", m["Answer"])  //elem = m[key] => 要素の取得 // => The value: 42

	m["Answer"] = 48  //要素の更新
	fmt.Println("The value:", m["Answer"])  // => The value: 48

	delete(m, "Answer")  //delete(m, key) => 要素の削除
	fmt.Println("The value", m["Answer"])  // => The value 0

	v, ok := m["Answer"]  //キーに対する要素が存在するかどうかの確認 => 2つ目の変数 ok が、存在する場合は true, しない場合は false となる
	fmt.Println("The value:", v, "Present?", ok)  // => The value: 0 Present? false
}