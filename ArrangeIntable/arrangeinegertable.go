// This program arranges a slice of int in ascending order.
package main

import "fmt"

func arrangeTable(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}

		}
	}
}

func main() {
	s := []int{5, 4, 3, 2, 1, 0}
	arrangeTable(s)
	fmt.Println(s)
}
