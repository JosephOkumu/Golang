package main

import (
	"fmt"
	"os"
)

func matchString(slice []string) {
	for _, v := range slice {
		if v == "galaxy 01" || v == "galaxy" || v == "01" {
			fmt.Println("Alert!!!")
		}
	}
	return

}

func main() {
	//This line takes the arguments from the second position, ignoring the program name
	args := os.Args[1:]

	matchString(args)

}
