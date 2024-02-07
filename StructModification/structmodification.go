package main

import "fmt"

// Define a struct type called Pilot with four fields: Name, Life, Age, and Aircraft
type Pilot struct {
	Name     string
	Life     float32
	Age      int
	Aircraft int
}

func main() {
	// Declare a variable donnie of type Pilot
	var donnie Pilot
	// Initialize the fields of the donnie variable
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = AIRCRAFT1

	// Print the donnie variable
	fmt.Println(donnie)
}

// Declare a constant named AIRCRAFT1 with the value 1
const AIRCRAFT1 = 1
