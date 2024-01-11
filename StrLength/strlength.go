//This program prints the length of a string.

package main

import "fmt"

func strLength(str string) int {
	return len(str)
}

func main() {

	l := strLength("Hello World!")
	fmt.Println(l)
}
