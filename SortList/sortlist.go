// This program sorts a shopping list

package main

import "fmt"

func sortList(slice []string) []string {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if len(slice[i]) > len(slice[j]) {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}

func main() {
	slice := []string{"Pineapple", "Honey", "Mushroom", "Tea", "Pepper", "Milk"}
	fmt.Println(sortList(slice))
}
