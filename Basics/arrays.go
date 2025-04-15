package main

import (
	"fmt"
)

func arr() {
	var a [5]int
	fmt.Println("emp: ", a) // by default array fills with 0 as elements

	a[4] = 100
	fmt.Println("set:", a) //changing values

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b) //declaring and initialization

	b = [...]int{100, 3: 400, 500} //basically after 3 put 400 and 500
	fmt.Println("idx:", b)

	var s []string
	s = make([]string, 3) //making slice of the string
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
}
