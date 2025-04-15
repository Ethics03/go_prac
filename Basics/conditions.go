package main

import (
	"fmt"
	"time"
)

func conditions() {

	var i int = 11

	if i < 10 {
		fmt.Printf("This is less than 10")
	} else if i > 10 {
		fmt.Printf("It is greater than 10")
	} else {
		fmt.Printf("It is 10")
	}

	switch i {
	case 1:
		fmt.Println("\nIt is one")
	case 2:
		fmt.Println("\nIt is 2")
	default:
		fmt.Println("\nIt is neither 1 nor 2")
	}

	//good use of switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	whatAmI := func(i interface{}) { //function declaration and pointing to the whatAmI variable
		switch t := i.(type) { // parameter is an interface
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
