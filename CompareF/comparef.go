// This program simulates the compare function
package main

import "fmt"

func compareF(a, b string) int {
	if a == b {
		return 0
	} else if a > b {
		return -1
	} else {
		return 1
	}
}

func main() {
	fmt.Println(compareF("Hello!", "Hello!"))
	fmt.Println(compareF("Salut!", "lut!"))
	fmt.Println(compareF("Ola!", "Ol"))

}
