package main

import "github.com/01-edu/z01"

func nthChar(s string, n int) rune {
	var char rune
	if n <= 0 || n > len(s) {
		return 0
	}
	for i, v := range s {
		if i == n-1 {
			char = v
			break
		}
	}
	return char
}

func main() {
	z01.PrintRune(nthChar("Hello!", 3))
	z01.PrintRune(nthChar("Salut!", 2))
	z01.PrintRune(nthChar("Bye!", -1))
	z01.PrintRune(nthChar("Bye!", 5))
	z01.PrintRune(nthChar("Ola!", 4))
	z01.PrintRune('\n')

}
