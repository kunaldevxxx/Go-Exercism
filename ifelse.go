package main

import "fmt"

func main() {
    var age int
    fmt.Print("Enter your age: ")
    fmt.Scan(&age)

    if age >= 18 {
        fmt.Println("eligible")
    } else {
        fmt.Println("not eligible")
    }
}