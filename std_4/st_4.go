package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("---INTERFACES---")
	// Interfaces are named collections of method signatures
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)

}

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

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 { // circle is a receiver of a method, passed by value
	return 2 * math.Pi * c.radius
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func measure(g geometry) {
	fmt.Println("geometry is ", g)
	fmt.Println("area is ", g.area())
	fmt.Println("perim is ", g.perim())
}
