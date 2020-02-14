package main

import "fmt"

func main() {
    var aValue float64 = 1.23
    var aPointer *float64 = &aValue
    fmt.Println("aPointer:", aPointer) // prints something like "aPointer: 0x1040a128"
    fmt.Println("*aPointer:", *aPointer) // prints "*aPointer: 1.23"
}