package main

import "fmt"

func optional() int {
        sum := 1
        for ; sum < 100 ; {
                sum += sum
        }
        return sum
}

func optional_omit() (sum int) {
        sum = 2
        for sum < 50 {
                sum += sum
        }
        return
}

func forever() {
        for {
        }
}

func main() {
        sum := 0
        for i:=0; i<10; i++ {
                sum += i
        }
        fmt.Println(sum)
        fmt.Println(optional())
        fmt.Println(optional_omit())
}
