package main

import (
    "fmt"
    "log"
    "math"
)

func main() {
    squareRoot, err := squareRoot(9)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(squareRoot)
}


//it can return multiple values from a function. 
//usually used to communicate errors.
func squareRoot(x float64) (float64, error) {
    if x < 0 {
        return 0, fmt.Errorf("can't take square root of negative number")
    }
    return math.Sqrt(x), nil
}