package main

import "fmt"

func main() {
    beforeIf := 888
    if true {
        //withinIf := 999
        fmt.Println("if")
    }
    fmt.Println(beforeIf) // OK
    //fmt.Println(withinIf) // Error!
}