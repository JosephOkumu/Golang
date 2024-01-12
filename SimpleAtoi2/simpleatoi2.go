// This program is a simulation of the atoi function. It tests for conversion for non-digit characters.
package main

import (
	"fmt"
)

func simpleAtoi2(str string) int {
	n := 0
	for _, char := range str {
		//checks if character is between 0 and 9
		if char >= '0' && char <= '9' {
			//The multiplication by 10 effectively shifts the existing digits to the left,
			//making room for the new digit on the right
			n = n*10 + int(char-'0')
		} else {
			//Returns 0 when it encounters a non-digit
			return 0
		}
	}
	return n
}

func main() {
	fmt.Println(simpleAtoi2("12345"))
	fmt.Println(simpleAtoi2("0000000012345"))
	fmt.Println(simpleAtoi2("012 345"))
	fmt.Println(simpleAtoi2("Hello World!"))
}
