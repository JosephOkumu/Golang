package main

import "fmt"

// The multiPointers function takes multiple pointers to integers as input and swaps their values.
func multiPointers(a ***int, b *int, c *******int, d ****int) {
	// Save the values of the pointers to temporary variables.
	tempA := ***a
	tempB := *b
	tempC := *******c
	tempD := ****d

	// Swap the values of the pointers.
	*******c = tempA
	****d = tempC
	*b = tempD
	***a = tempB
}

func main() {
	// Create multiple pointers to integers.
	x := 5
	y := &x
	z := &y
	a := &z

	w := 2
	b := &w

	u := 7
	e := &u
	f := &e
	g := &f
	h := &g
	i := &h
	j := &i
	c := &j

	k := 6
	l := &k
	m := &l
	n := &m
	d := &n

	// Print the initial values of the pointers.
	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)

	// Call the multiPointers function to swap the values of the pointers.
	multiPointers(a, b, c, d)

	// Print the final values of the pointers.
	fmt.Println("After using multipointers")
	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)
}
