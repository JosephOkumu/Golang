// This program is of a bot that can hold a simple greeting
package main

import "fmt"

func main() {
	fmt.Println("In one word, how do you do today?")
	var response1 string

	fmt.Scan(&response1)

	if response1 == "good" || response1 == "amazing" || response1 == "great" {
		fmt.Println("Glad to hear, enjoy yourself!")
	} else {
		fmt.Println("Sorry to hear that. Hope things get well!")
	}
}
