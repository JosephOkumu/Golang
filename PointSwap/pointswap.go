// This program swaps pointers.
package main

import "fmt"

func pointSwap(a, b *int) {
	*a, *b = *b, *a

}

func main() {
	a := 0
	b := 1

	pointSwap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
