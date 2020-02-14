package main

import "fmt"

func main(){
	myNumber := 2.6
	halve(myNumber)
	//Does nothing!
	//Print 2.6!
	fmt.Println("myNumber in 'main':", myNumber)
}

func halve( number float64 ){
	number = number /2
	//Print 1.3, but change is only effective here!
	fmt.Println("number in 'halve' :", number)
}