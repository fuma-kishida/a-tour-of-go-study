package main

import "fmt"

func main() {
	// := や var = 明示的に型を指定せずに変数を宣言する場合、
	//変数の型は右側の変数から型推論される
	v := 0.867 + 0.5i  
	fmt.Printf("v is of type %T\n", v)
}