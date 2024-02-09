package main

import "fmt"

// Define a function to process the input string and return a modified string
func fivePacks(str string) string {
	// If the input string is empty, return a newline
	if str == "" {
		return "\n"
	}
	// If the length of the input string is less than 5, return "Invalid Output" followed by a newline
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	// Initialize an empty string to store the modified output
	score := ""
	// Initialize a counter to keep track of the number of characters processed
	cnt := 0

	// Iterate over each character in the input string
	for i, v := range str {
		// If the character is not a space and the character count is less than 5, add the character to the output string
		if v != ' ' && cnt != 5 {
			score += string(v)
			cnt++
		} else if cnt == 5 {
			// If the character count reaches 5, add a space to the output string and reset the character count
			score += " "
			cnt = 0
		}
		// If it's the last character and the output string is not empty and ends with a space, remove the trailing space
		if i == len(str)-1 && len(score) > 0 && score[len(score)-1] == ' ' {
			score = score[:len(score)-1]
		}
	}
	// Add a newline to the end of the modified output string and return it
	score += "\n"
	return score
}

func main() {
	// Call the fivePacks function with different input strings and print the output
	fmt.Print(fivePacks("deliciousbread"))
	fmt.Print(fivePacks("This is a loaf of bread"))
	fmt.Print(fivePacks("loaf"))
}
