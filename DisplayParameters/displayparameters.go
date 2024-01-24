package main

import (
	"os"

	"github.com/01-edu/z01"
)

func displayParameters() {
	args := os.Args[1:]

	for _, v := range args {
		for _, j := range v {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}

func main() {
	displayParameters()
}
