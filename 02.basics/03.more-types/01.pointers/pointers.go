package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i  //i のメモリ上のアドレス
	fmt.Println(p)  //0x40e020
	fmt.Println(*p)  //ポインタによって i を引き出す //42
	fmt.Println(i)   //42

	p = &j
	*p = *p / 37
	fmt.Println(j)  //73
}