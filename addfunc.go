//This program adds two arguments and displays the result.

package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(23, 38))
}
