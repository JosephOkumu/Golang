package main

import (
	"fmt"
)

// The skip function takes a string as input and returns every third character as a string.
func skip(str string) string {
	var wrd string

	// If the input string is empty or has less than 3 characters, return a newline character.
	if len(str) < 3 {
		return "\n"
	}

	// Iterate through the input string and add every third character to the output string.
	for i := 2; i < len(str); i += 3 {
		wrd += string(str[i])
	}

	// Return the output string with a newline character.
	return wrd + "\n"
}

func main() {
	// Call the skip function with different input strings and print the output.
	fmt.Print(skip("1010101010"))
	fmt.Print(skip(""))
	fmt.Print(skip("t w e l v e"))
	fmt.Print(skip("12"))
}
