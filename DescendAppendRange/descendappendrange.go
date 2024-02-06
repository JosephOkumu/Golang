package main

import "fmt"

func descAppendRange(max, min int) []int {
	var slice []int

	if max <= min {
		return slice
	} else {
		for i := max; i >= min; i-- {
			slice = append(slice, i)

		}
	}
	return slice[:len(slice)-1]
}

func main() {
	fmt.Println(descAppendRange(10, 5))
	fmt.Println(descAppendRange(5, 10))
}
