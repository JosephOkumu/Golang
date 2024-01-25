package main

import (
	"os"

	"github.com/01-edu/z01"
)

func arrangeParamet() {
	args := os.Args[1:]

	for i := 0; i < len(args)-1; i++ {
		for j := i + 1; j < len(args); j++ {
			if args[i] > args[j] {
				args[i], args[j] = args[j], args[i]
			}
		}
	}
	for _, arg := range args {
		for _, char := range arg {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')

	}

}

func main() {
	arrangeParamet()
}
