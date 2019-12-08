//Go のプログラムは package で構成される 
//プログラムは main パッケージから開始される
package main

import (
	"fmt"  // "fmt" パッケージを import
	"math/rand"  //"math/rand" パッケージを import
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}