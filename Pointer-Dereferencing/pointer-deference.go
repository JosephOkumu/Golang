//This code illustrates the use of pointers and deferencing concept.

package main

import "fmt"

func main() {
	star := "Polaris"
	//The line of code below creates a pointer
	//The two lines below can also be combined to one as starAddress := &star
	var starAddress *string
	starAddress = &star

	//This line below is called dereferencing. Which means changing the value in that address
	*starAddress = "Sirius"

	fmt.Println("The actual value of star is", star)
	fmt.Println(starAddress)
}
