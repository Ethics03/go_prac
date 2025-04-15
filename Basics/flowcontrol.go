package main

import (
	"fmt"
)

func flow() {
	// golang dosent have a while loop its just the same
	//thing in for with conditions and incrementations
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 { //new way of setting range (0-(n-1))
		fmt.Println(i)
	}

	//break and continue are same as cpp
}
