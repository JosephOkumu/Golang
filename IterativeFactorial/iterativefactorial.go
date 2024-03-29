// This program fincs the factorial of a number using for. it checks for errors/overflows
package main

import "fmt"

func iterativeFactorial(nb int) int {
	result := 1
	if nb < 0 || nb > 63 {
		return 0
	}
	for i := 1; i <= nb; i++ {
		result = result * i
	}
	return result
}

func main() {
	arg := 64
	fmt.Println(iterativeFactorial(arg))

}
