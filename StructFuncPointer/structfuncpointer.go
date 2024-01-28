package main

import "fmt"

type Triangle struct {
	height float32
	base   float32
}

//The code below takes a struct Triangle with a pointer and passes it to a function updateBase
//the function updateBase changes the base of the Triangle of the instance triangle

func (triangle *Triangle) updateBase(newBase float32) {
	triangle.base = newBase
}

func (triangle Triangle) area() float32 {
	return (triangle.base * triangle.height) / 2
}

func main() {
	triangle := Triangle{10, 4}
	//The code below passes the struct to the function updatebase and gives it a value of 13
	triangle.updateBase(13)
	fmt.Println(triangle)
	//The code below prints the area of the triangle.
	fmt.Println(triangle.area())
}
