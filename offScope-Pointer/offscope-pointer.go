//This program illustrates how to change a variable that is outside the scope using a pointer.

package main

import "fmt"

func brainwash(saying *string) {
	// Dereferencing/changing the value in the address of saying to beep boop.
	*saying = "Beep Boop."

}

func main() {
	greeting := "Hello there!"

	// the value in the address of greeting is going to be changed to beep boop when we call brainwash
	brainwash(&greeting)

	fmt.Println("greeting is now:", greeting)
}
