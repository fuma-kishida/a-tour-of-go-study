package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := float64(2.)
    s := float64(0)
    for i := 0; i < 10; i++ {
        z = z - (z*z - x)/(2*z)
        if math.Abs(z - s) < 1e-10 {
            break;
        }
        s = z
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(math.Sqrt(2))
}