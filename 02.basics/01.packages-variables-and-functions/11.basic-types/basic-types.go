package main

import (
	"fmt"
	"math/cmplx"
)

//Go の 基本型たち（一例）
//インポートステートメント同様、変数はまとめて宣言が可能
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 -1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}