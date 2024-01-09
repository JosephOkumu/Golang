// This program displays on a single line, the numbers entered.
package main

import "github.com/01-edu/z01"

func dispDigit(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	} else if n < 0 {
		//The code below prints negative.
		z01.PrintRune('-')
		//The line of code below makes the negative n above positive.
		n = -n
	}
	//The printNba function in the line below is called using the positive n above
	printNba(n)
}

func printNba(n int) {
	if n/10 != 0 {
		//The line below calls itself recursively if n/10 != 0
		printNba(n / 10)
	}
	z01.PrintRune('0' + rune(n%10))
}

func main() {
	dispDigit(-123)
	dispDigit(0)
	dispDigit(123)
}
