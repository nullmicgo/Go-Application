package main

import (
    "fmt"
    "os"
)

func main() {
    fileInfo, error := os.Stat("existent.txt")
    if error != nil {
        fmt.Println(error)
    } else {
        fmt.Println(fileInfo.Size())
    }
    fileInfo, error = os.Stat("nonexistent.txt")
    if error != nil {
        fmt.Println(error)
    } else {
        fmt.Println(fileInfo.Size())
    }
}