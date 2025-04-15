package main

import (
	"fmt"
	"maps"
)

func mapz() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 3}
	if maps.Equal(n, n2) { // maps.Equal() -> for checking the equality of the two maps
		fmt.Println("n == n2")
	} else {
		fmt.Println("n != n2")
	}

}
