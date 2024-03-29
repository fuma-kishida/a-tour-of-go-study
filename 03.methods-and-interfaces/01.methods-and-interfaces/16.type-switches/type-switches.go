package main

import "fmt"

func do(i interface{}) {
	switch v:= i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)  //21 について swtch-case => Twice 21 is 42
	do("hello")  //"hello" について swtch-case => "hello" is 5 bytes long
	do(true)  //true について swtch-case => I don't know about type bool!
}