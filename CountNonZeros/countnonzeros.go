// This program return the number of elements with non-zero value
// It then deletes the elements with zero value in the slice.
package main

import "fmt"

func countNotZero(ptr *[]string) int {
	s := *ptr // Dereferences the pointer to get the slice
	j := 0

	for i := 0; i < len(s); i++ {
		if s[i] != "" { // Check if the element is not an empty string
			s[j] = s[i] // Move the non-empty element to the front of the slice
			j++         // Increment the count of non-empty elements
		}
	}
	*ptr = s[:j]
	return j
}

func main() {
	a := make([]string, 6)
	a[0] = "a"
	a[2] = "b"
	a[4] = "c"

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("Size after deleting:", countNotZero(&a))

	for _, v := range a {
		fmt.Println(v)
	}
}
