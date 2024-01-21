// This program prints the first rune.
package main

import "github.com/01-edu/z01"

func initialRune(s string) rune {
	list := []rune(s)
	return list[0]

}

func main() {
	z01.PrintRune(initialRune("Gopher"))
}
