package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")  //新しいgoroutine(並行に動作している関数)の実行
	say("hello")
}  

/*
=>
world
hello
hello
world
world
hello
hello
world
world
hello
*/
