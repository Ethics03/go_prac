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
	fmt.Println("uninit:", s, s == nil, len(s) == 0)
	s = make([]string, 3) //making slice of the string
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b) //  ... is basically letting the compiler decide the size of the array

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)
	
	whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")

}
