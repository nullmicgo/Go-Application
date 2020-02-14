package main

import (
    "fmt"
)

func main() {
    fmt.Print("You win: ")
    doorNumber := 2
    switch doorNumber {
    case 1:
        fmt.Println("a new car!") // not printed
    case 2:
        fmt.Println("a llama!") // printed
    default:
        fmt.Println("a goat!") // not printed
    }
}