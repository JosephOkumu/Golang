package main

import (
	"os"

	"github.com/01-edu/z01"
)

func revDisplayedparams() {
	args := os.Args[1:]
	str := ""

	for i := len(args) - 1; i >= 0; i-- {
		word := args[i]
		str = str + " " + word
		//fmt.Println(str)
	}
	for _, v := range str {
		z01.PrintRune(v)
		if v == ' ' {
			z01.PrintRune('\n')
		}

	}

}

func main() {
	revDisplayedparams()
}
