// This is iterative power. The program finds the value of a number nb, raised to power, power.
package main

import "fmt"

func iterPower(nb, power int) int {
	result := 1
	if power < 0 {
		return 0
	}
	for i := 0; i < power; i++ {
		result = result * nb
	}
	return result
}

func main() {
	fmt.Println(iterPower(2, 4))
}
