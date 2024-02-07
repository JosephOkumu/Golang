//This program displays the correct player position

package main

import "fmt"

// playerPosition takes a 2D slice of strings and sorts it based on the second element of each inner slice in ascending order
func playerPosition(player [][]string) [][]string {
	// Iterate over the player slice
	for i := 0; i < len(player)-1; i++ {
		// Iterate over the player slice starting from the next element after i
		for j := i + 1; j < len(player); j++ {
			// Check if the length of the inner slices is at least 2 and if the second element of player[i] is greater than the second element of player[j]
			if len(player[i]) >= 2 && len(player[j]) >= 2 && player[i][1] > player[j][1] {
				// Swap the positions of player[i] and player[j] in the slice
				player[i], player[j] = player[j], player[i]
			}
		}
	}
	// Return the sorted player slice
	return player
}

func main() {
	// Create slices representing the positions
	p4 := []string{"4th Place"}
	p3 := []string{"3rd Place"}
	p2 := []string{"2nd Place"}
	p1 := []string{"1st Place"}

	// Construct a 2D slice of strings
	position := [][]string{p4, p3, p2, p1}
	// Call the playerPosition function with the position slice and print the result
	fmt.Println(playerPosition(position))
}
