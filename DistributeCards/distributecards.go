package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

// The distCards function takes a slice of integers representing a deck of cards and prints the cards for each player.
func distCards(deck []int) {
	// Iterate through each player
	for i := 0; i < 4; i++ {
		// Print the player number
		fmt.Printf("Player %d: ", i+1)

		// Iterate through the cards for the current player
		for j := i * 3; j < (i+1)*3; j++ {
			// Print the card number
			fmt.Printf("%d", deck[j])

			// If it's not the last card for the player, print a comma and a space
			if j != (i+1)*3-1 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		// Print a new line after printing all the cards for the player
		z01.PrintRune('\n')
	}
}

func main() {
	// Create a deck of cards
	deck := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	// Call the distCards function to print the cards for each player
	distCards(deck)
}
