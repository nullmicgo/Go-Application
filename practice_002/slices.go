package main

import "fmt"

func main() {
	//Slices
	a := [5]int{0, 1, 2, 3, 4}
	s1 := a[0:3]
	s2 := a[2:5]
	fmt.Println(a, s1, s2)

	a[2] = 88
	fmt.Println(a, s1, s2)

	s2[0] = 999
	fmt.Println(a, s1, s2)

	s1 = s1[0:4]

	fmt.Println("len(s1):", len(s1), "cap(s1):", cap(s1))
	fmt.Println("len(s2):", len(s2), "cap(s2):", cap(s2))

	s2 = append(s2, 5)
	fmt.Println(a, s1, s2)

	fmt.Println("len(s1):", len(s1), "cap(s1):", cap(s1))
	fmt.Println("len(s2):", len(s2), "cap(s2):", cap(s2))

	s := []int{1, 2, 3}
	s = append(s, 4, 5)
	s = append(s, 6, 7, 8)
	fmt.Println(s)

	st := []int{1, 2, 3}
	for i, v := range st {
		fmt.Println("element:", i, "value:", v)
	}

}
