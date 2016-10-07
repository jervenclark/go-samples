package main

import (
        "fmt"
        "time"
        "runtime"
)

func os() {
        fmt.Print("Go runs on ")
        switch os := runtime.GOOS; os {
        case "darwin":
                fmt.Println("OSX")

        case "linux":
                fmt.Println("Linux")

        default:
                fmt.Printf("%s.", os)
        }
}

func order() {
        fmt.Println("When's Saturday?")
        today := time.Now().Weekday()
        switch time.Saturday {
        case today + 0:
                fmt.Println("Today.")
        case today + 1:
                fmt.Println("Tomorrow.")
        case today + 2:
                fmt.Println("In two days.")
        default:
                fmt.Println("Too far away.")
        }
}

func no_condition() {
        t := time.Now()
        switch {
        case t.Hour() < 12:
                fmt.Println("Good Morning!")
        case t.Hour() < 17:
                fmt.Println("Good Afternoon!")
        default:
                fmt.Println("Good Evening!")
        }
}

func main() {
        os()
        order()
        no_condition()
}
