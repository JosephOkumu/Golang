// This program counts the letters in a string.
package main

import "fmt"

func letterCount(s string) int {
	count := 0
	for _, v := range s {
		if (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {
			count++

		}

	}
	return count
}

func main() {
	s := "Hello 34 World!    482 "
	n := letterCount(s)
	fmt.Println(n)

}
