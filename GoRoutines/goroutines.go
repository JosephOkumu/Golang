package main

import (
	"fmt"
	"time"
)

// Function to print a string s five times with a 100 millisecond delay between each print
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// Function to demonstrate the usage of goroutines to execute two functions concurrently
func main() {
	// Create a new goroutine to execute the say function with the argument "world"
	go say("world")

	// Execute the say function with the argument "Hello" in the main goroutine
	say("Hello")
}
