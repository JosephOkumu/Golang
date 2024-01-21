package main

import "github.com/01-edu/z01"

func lastChar(s string) rune {
	slice := []rune(s)

	return slice[len(slice)-1]
}

func main() {
	z01.PrintRune(lastChar("My Bestfriend"))
}
