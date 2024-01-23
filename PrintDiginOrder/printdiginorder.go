package main

import "github.com/01-edu/z01"

func printdiginOrder(n int) {

	if n < 0 {
		// Print nothing for negative numbers
		return
	} else if n == 0 {
		//Print zero if it is zero
		z01.PrintRune('0')
	}

	// Create an array to store counts of each digit
	counts := [10]int{}

	// Count the occurrences of each digit in the number
	for n > 0 {
		//finds the last digint in the number.
		digit := n % 10
		//Adds the last digit to count array
		counts[digit]++
		//updates n by removing the last digit.
		n /= 10
	}

	// Print the digits in ascending order
	for i := 0; i < 10; i++ {
		for counts[i] > 0 {
			z01.PrintRune(rune('0' + i))
			counts[i]--
		}
	}
}

func main() {
	printdiginOrder(321)
	printdiginOrder(0)
	printdiginOrder(321)
}
