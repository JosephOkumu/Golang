package main

import (
	"os"

	"github.com/01-edu/z01"
)

func arrangeParamet() {
	args := os.Args[1:]
	//str := ""

	for i := 0; i < len(args); i++ {
		//wrd := args[i]
		for j := 0; j < len(args)-1; j++ {
			for k := j + 1; k < len(args); k++ {
				if args[j] > args[k] {
					args[j], args[k] = args[k], args[j]
				}

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
