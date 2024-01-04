//this program checks if a number is positive or not.
//Prints T if the number is positive and prints F if the number is not positive.

package main

import (
	"github.com/01-edu/z01"
)

func isPositive(nb int) {
	if nb > 0 {
		z01.PrintRune('T')
	} else if nb < 0 {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}

func main() {
	isPositive(1)
	isPositive(-3)

}
