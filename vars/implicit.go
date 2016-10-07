package main

import "fmt"

// := construct is not available here

func main() {
        var i, j int = 1, 2
        // := is available here
        k := 3
        c, python, java := true, false, "no!"
        fmt.Println(i, j, k, c, python, java)
}
