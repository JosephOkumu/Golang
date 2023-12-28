// This program swaps two strings and displays the result after the swap.
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Hello", "world")
	fmt.Println(a, b)
}
