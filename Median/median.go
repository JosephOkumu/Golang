package main

import "fmt"

func median(a, b, c, d, e int) int {
	slice := []int{a, b, c, d, e}
	mid := (len(slice) / 2)

	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice[mid]
}

func main() {
	middle := median(2, 3, 8, 5, 7)
	fmt.Println(middle)
}
