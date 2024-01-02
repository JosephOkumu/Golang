package main

import "fmt"

//The fmt.Scan is for capturing inputs.
//fmt.Scan expects addresses for its arguments. That is why we put & before the argument

func main() {
	fmt.Println("What would you like for lunch?")

	var food string
	var drink string
	//scan only takes one argument only
	fmt.Scan(&food)

	fmt.Println("Any drink?")
	fmt.Scan(&drink)

	fmt.Printf("Sure, we can have %v for lunch and some %v ", food, drink)
}
