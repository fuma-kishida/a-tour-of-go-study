package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
    "fmt"
)

func WordCount(s string) map[string]int {
    // I am learning Go! -> [I am learning Go!]
    // func Fields(s string) []string
    sf := strings.Fields(s)
    // Go!
    fmt.Println(sf[3])
    // 4
    num := len(sf)
    ret := make(map[string]int)
    /*
    map[I:1]
    map[I:1 am:1]
    map[I:1 am:1 learning:1]
    map[learning:1 Go!:1 I:1 am:1]
    */
    for i := 0; i < num; i++ {
        (ret[sf[i]])++
        fmt.Println(ret)
    }

    return ret
}

func main() {
    wc.Test(WordCount)
}