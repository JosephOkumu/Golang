// The program returns a slice of int where min is included and max is not included.
package main

import "fmt"

func appendSlice(min, max int) []int {
	var slice []int

	for i := min - 1; i < max-1; i++ {
		slice = append(slice, i+1)
	}
	return slice
}

func main() {
	fmt.Println(appendSlice(20, 10))
}
