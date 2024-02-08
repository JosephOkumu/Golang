package main

import "github.com/01-edu/z01"

// Define a struct type Door with a state field of type string
type Door struct {
	state string
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

// Define a function to print a string and update the door state to "OPEN"
func openDoor(ptrDoor *Door) bool {
	printStr("Door opening...")
	ptrDoor.state = "OPEN"
	return true
}

// Define a function to print a string and check if the door state is "OPEN"
func isDoorOpen(door Door) bool {
	printStr("Is the Door opened?")
	return door.state == "OPEN"
}

// Define a function to print a string and check if the door state is "CLOSE"
func isDoorClose(ptrDoor *Door) bool {
	printStr("Is the door closed?")
	return ptrDoor.state == "CLOSE"
}

// Define a function to print a string and update the door state to "CLOSE"
func closeDoor(ptrDoor *Door) bool {
	printStr("Door closing...")
	ptrDoor.state = "CLOSE"
	return true
}

func main() {
	// Create a new Door instance
	door := &Door{}

	// Open the door
	openDoor(door)

	// If the door is not closed, open it
	if isDoorClose(door) {
		openDoor(door)
	}
	// If the door is open, close it
	if isDoorOpen(*door) {
		closeDoor(door)
	}
	// If the door is still open, close it
	if door.state == "OPEN" {
		closeDoor(door)
	}
}
