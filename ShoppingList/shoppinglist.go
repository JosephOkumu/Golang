package main

import "fmt"

func shoppingListCounter(str string) map[string]int {
	summary := make(map[string]int) // Create a map to store the count of each item
	item := ""                      // Initialize an empty string to store each item

	for _, char := range str { // Iterate through each character in the input string
		if char != ' ' { // If the character is not a space, add it to the current item
			item += string(char)
		} else { // If the character is a space, it means the current item is complete
			summary[item]++ // Increment the count of the current item in the map
			item = ""       // Reset the current item to an empty string for the next item
		}
	}
	summary[item]++ // Increment the count of the last item in the map
	return summary  // Return the map containing the count of each item
}

func main() {
	summary := "Burger Water Carrot Coffee Water Water Chips Carrot Carrot Burger Carrot Water"
	for index, element := range shoppingListCounter(summary) { // Call the shoppingListCounter function and iterate through the returned map
		fmt.Println(index, "=>", element) // Print each item and its count
	}
}
