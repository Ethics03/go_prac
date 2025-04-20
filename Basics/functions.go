package main

import "fmt"

func functions() {
	a := 2
	b := 3
	fmt.Println(plus(a, b))
	fmt.Println(show2(a, b))

	sum(1, 2, 3, 4, 5)
}

func plus(a int, b int) int {
	return a + b
}

func show2(a int, b int) (int, int) {
	return a, b // MULTIPLE RETURN VALUE function
}

func sum(nums ...int) { //VARIADIC  functions
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
