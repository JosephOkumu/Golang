package main

import "fmt"

// This function takes as arguments a function and a slice of integer as parameters.
func toEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}

func PrintNbr(n int) {
	fmt.Print(n)
}
func main() {
	a := []int{1, 2, 3, 4, 5, 6}

	//to the function toEach, we pass a function PrintNbr and a slice, a
	toEach(PrintNbr, a)
}
