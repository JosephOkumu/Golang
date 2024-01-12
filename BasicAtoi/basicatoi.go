// This program simulates the behaviour of the Atoi function.
package main

import (
	"fmt"
)

func BasicAtoi(s string) int {
	n := 0
	for _, v := range s {
		//where 48 is the ascii value of char 0
		n = n*10 + int(v-48)
	}
	return n
}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000"))
	fmt.Println(BasicAtoi("00001234"))

}
