package main

import "github.com/01-edu/z01"

// declaring a struct called point
type point struct {
	x int
	y int
}

// creating a pointer ptr that points to the point struct
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

// This prints the characters one by one
func iStr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func isInt(n int) {
	val := '0'
	//This section checks whether n is greater than 0 and calls the IsInt function recursively if so.
	if n/10 > 0 {
		isInt(n / 10)
	}
	//When n is !> 0 it iterates from 0 to that value of n, in our case 4
	//It then goes back to the first recursive call and ginds the remainder; in this case 42%10 == 2
	for i := 0; i < n%10; i++ {
		val++
	}
	//This prints 4 first then 2
	z01.PrintRune(val)

}

func main() {
	//This declares points as an empty point struct
	points := point{}
	//The declared empty struct is passed to the setpoint function call and receives the values of x and y
	setPoint(&points)
	//This prints x first
	iStr("x = ")
	//prints the value when 42 is passed
	isInt(points.x)
	//Prints y after printing x.
	iStr(", y = ")
	//prints the value when 21 is passed
	isInt(points.y)
	z01.PrintRune('\n')

}
