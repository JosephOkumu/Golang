// This program displays on a single line, the numbers entered.
package main

import "github.com/01-edu/z01"

func dispDigit(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	} else if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	printNba(n)
}

func printNba(n int) {
	if n/10 != 0 {
		printNba(n / 10)
	}
	z01.PrintRune('0' + rune(n%10))
}

func main() {
	dispDigit(-123)
	dispDigit(0)
	dispDigit(123)
}
