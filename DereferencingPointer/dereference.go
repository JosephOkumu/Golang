/*This program divides the dereferenced value of a by the dereferenced value of b.
Stores the result of the division in the int which a points to.
Stores the remainder of the division in the int which b points to.*/

package main

import "fmt"

func dereference(a, b *int) {
	i := *a
	j := *b

	*a = i / j
	*b = i % j
}

func main() {
	a := 13
	b := 2

	dereference(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
