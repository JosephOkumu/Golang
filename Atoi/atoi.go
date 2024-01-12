// Simulation of the Atoi function with no error detection.
package main

import "fmt"

func Atoi(s string) int {
	num := 0
	sign := 1
	for i, v := range s {
		if i == 0 && v == '-' {
			sign = -1
		} else if i == 0 && v == '+' {
			sign = 1
		} else if v >= '0' && v <= '9' {
			num = num*10 + int(v-48)
		} else {
			return 0
		}
	}
	return num * sign
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}
