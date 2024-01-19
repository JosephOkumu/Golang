//This program finds the square root of a number.

package main

import "fmt"

func root(n int) int {
	for i := 0; i < n; i++ {
		if i*i == n {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(root(625))
}
