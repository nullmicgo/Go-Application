package main

import "fmt"

func main() {
	car := Car {
		Doors: 4,
		Transmission: "automatic",
		Odometer: 36000,
	}
	describe(&car)
}

func describe(c *Car){
	fmt.Println("A", c.Doors, "door")
	fmt.Println(c.Transmission,"car")
	fmt.Println("with", c.Odometer, "miles")
	
	
}