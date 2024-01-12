// This program simulates the behaviour of the Atoi function.
package main

import (
	"fmt"
	"strconv"
)

func BasicAtoi(s string) int {
	n := 0
	if v, err := strconv.Atoi(s); err == nil {
		n = n + v
	}
	return n
}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000"))
	fmt.Println(BasicAtoi("00001234"))

}
