package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

type square struct {
	side float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (r square) area() float64 {
	return r.side * r.side
}

func (r square) perim() float64 {
	return r.side * 4
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func interfacePrac() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	s := square{side: 6}

	measure(r)
	measure(c)
	measure(s)

	detectCircle(r)
	detectCircle(c)
}
