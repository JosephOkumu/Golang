/*
	This is a program for choosing a random no:

We can import math/rand and use rand.Intn(number) but this will always give the same number.
So we import time to make random nos.
N/B :- In this version of go rand.Intn(number) just works fine.
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//The line below of rand.Seed... Must be included otherwise the code prints same number
	//rand.Seed(time.Now().UnixNano())
	// The line below must be used with the line above, otherwise the code prints same number throughout.
	amountLeft := rand.Intn(10000)

	fmt.Println("amountLeft is: ", amountLeft)

	if amountLeft > 5000 {
		fmt.Println("What should I spend this on?")
	} else {
		fmt.Println("Where did all my money go?")
	}
}
