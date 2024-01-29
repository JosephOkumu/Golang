package main

import "github.com/01-edu/z01"

func printSliceonline(sl []string) {
	for _, v := range sl {
		for _, char := range v {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}

}

func main() {
	test := []string{"Hello", "how", "are", "you"}
	printSliceonline(test)
}
