// This program uses recursion to find the value of number raised to a power
package main

import "fmt"

func recursePower(nb, power int) int {
	if power < 0 {
		return 0
	} else if power == 1 {
		return nb
	} else {
		return nb * recursePower(nb, power-1)
	}
}

func main() {
	fmt.Println(recursePower(4, 3))
}
