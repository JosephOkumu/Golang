// This program reverses the order of the menu
package main

import "fmt"

// reverseMenu takes a slice of strings and returns a new slice with the elements in reverse order
func reverseMenu(menu []string) []string {
	// Create a new slice with the same length as the input slice
	slice := make([]string, len(menu))
	// Iterate over the input slice in reverse order
	for i := len(menu) - 1; i >= 0; i-- {
		// Assign each element of the input slice to the new slice in the opposite order
		slice[len(menu)-1-i] = menu[i]
	}
	// Return the new slice with the elements in reverse order
	return slice
}

func main() {
	// Call the reverseMenu function with a sample menu and print the result
	fmt.Println(reverseMenu([]string{"desserts", "mains", "drinks", "starters"}))
}
