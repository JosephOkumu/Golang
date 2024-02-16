package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)  // Create a channel that sends tick events every 100 milliseconds
	boom := time.After(500 * time.Millisecond) // Create a channel that sends a boom event after 500 milliseconds

	for {
		select {
		case <-tick: // If a tick event is received, print "tick."
			fmt.Println("tick.")
		case <-boom: // If a boom event is received, print "BOOM!" and exit the program
			fmt.Println("BOOM!")
			return
		default: // If neither a tick nor a boom event is received, print a dot and wait for 50 milliseconds
			fmt.Println("   .")
			time.Sleep(50 * time.Millisecond)
		}
	}

}
