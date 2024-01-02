// This program highlights the different usage of printf verbs
package main

import "fmt"

func main() {
	floatExample := 1.75
	// %T displays the type
	fmt.Printf("Working with a %T", floatExample)

	fmt.Println("\n***") // Added for spacing

	yearsOfExp := 3
	reqYearsExp := 15
	// %d displays the digit/number. Note the first argument takes the first % verb.
	fmt.Printf("I have %d years of Go experience and this job is asking for %d years.", yearsOfExp, reqYearsExp)

	fmt.Println("\n***") // Added for spacing

	stockPrice := 3.50
	// %f displays float values. %.2f specifies the decimal places.
	fmt.Printf("Each share of Gopher feed is $%.2f!", stockPrice)

	fmt.Println("\n***")

	word := "C"

	fmt.Printf("She said the %v word", word)
}
