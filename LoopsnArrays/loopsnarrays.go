package main

import (
	"fmt"
)

func main() {

	menu := []string{"Hamburgers", "Cheeseburgers", "Fries"}

	fmt.Println("The menu:")

	for i, v := range menu {
		fmt.Println("number", i, "item", v)
	}

	orders := map[string]string{
		"John":   "Cheeseburgers",
		"Janet":  "Hamburgers",
		"Jordan": "Fries",
	}

	orders["James"] = "Chicken Sandwich"

	fmt.Println("\nThe friend's orders:")

	for k, v := range orders {
		fmt.Println("Name:", k, "order-", v)
	}

}
