package main

import (
	"os"

	"github.com/01-edu/z01"
)

func displayProgname() {
	args := os.Args[0]
	for i := 0; i < len(args); i++ {
		ltr := rune(args[i])
		z01.PrintRune(ltr)
	}
}

func main() {
	displayProgname()
}
