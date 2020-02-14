package main

import "fmt"

func main() {
  ages := map[string]float64{}
  // Like the array/slice syntax, but you can use any value of the type you specified for the keys.
  ages["Alice"] = 12
  ages["Bob"] = 9
  fmt.Println(ages)
}


func HalfPriceSale(prices map[string]float64) map[string]float64 {
    halfPriced := make(map[string]float64)
    for key, value := range prices {
        halfPriced[key] = value / 2
    }
    return halfPriced
}