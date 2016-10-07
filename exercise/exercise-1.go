package main

import (
        "fmt"
        "math"
)

func sqrt(x float64) (float64) {
        z := x;
        i := 1
        for i<10 {
                z = z - ((math.Pow(z, 2) - x) / (2 *z))
                i++
        }
        return z
}

func main() {
        for i:=5; i<10; i++ {
                a, b := sqrt(float64(i*3)), math.Sqrt(float64(i*3))
                c := a -b
                fmt.Println(
                        a,
                        b,
                        c,
                )
        }
}
