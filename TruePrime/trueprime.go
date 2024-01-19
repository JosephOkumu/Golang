package main

import "fmt"

func truePrime(nb int) bool {

	for i := 2; i < nb; i++ {
		if nb <= 1 {
			return false
		} else if nb%i == 0 {
			return false
		}
	}
	return true

}

func main() {
	fmt.Println(truePrime(15))
}
