package main

import "fmt"

func pointDivmod(a int, b int, div *int, mod *int) {
	*div = a / b
	*mod = a % b

}

func main() {
	a := 13
	b := 2

	var div int
	var mod int

	pointDivmod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}
