package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func structs() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(newPerson("Jon"))

	cat := struct {
		name   string
		isCute bool
	}{
		"Rex",
		true,
	}

	fmt.Println(cat)
}
