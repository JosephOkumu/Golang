package main

import "fmt"

func recursFactorial(nb int) int {
	if nb < 0 || nb > 63 {
		return 0
	} else if nb == 0 || nb == 1 {
		return 1
	} else {
		return nb * recursFactorial(nb-1)
	}
}

func main() {
	arg := 6
	fmt.Println(recursFactorial(arg))
}
