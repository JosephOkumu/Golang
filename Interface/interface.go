package main

import (
	"fmt"
	"math"
)

// Define an interface called shape with a single method called area
type shape interface {
	area() float64
}

// Define a struct type called circle with a radius field
type circle struct {
	radius float64
}

// Define a struct type called rect with width and height fields
type rect struct {
	width  float64
	height float64
}

// Define a method for the rect struct to calculate its area
func (r rect) area() float64 {
	return r.width * r.height
}

// Define a method for the circle struct to calculate its area
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	// Create a circle and a rectangle instance
	c1 := circle{4.5}
	r1 := rect{5, 7}

	// Create a slice of shape interface and add the circle and rectangle instances to it
	shapes := []shape{c1, r1}

	// Iterate over each shape in the slice and print its area
	for _, shape := range shapes {
		fmt.Println(shape.area())
	}
}
