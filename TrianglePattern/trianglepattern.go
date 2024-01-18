package main

import "github.com/01-edu/z01"

func Triangle(x, y int) {
	for row := 0; row < y; row++ {
		for col := 0; col <= row; col++ {
			z01.PrintRune('*')
		}
		z01.PrintRune('\n')
	}
}

func main() {
	Triangle(5, 5)
}
