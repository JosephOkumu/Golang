package main

import "fmt"

func makeSlice(min, max int) []int {
	if max <= min {
		return nil
	}
	slice := make([]int, (max - min))

	for i := 0; i < max-min; i++ {
		slice[i] = min + i
	}

	return slice
}

func main() {
	fmt.Println(makeSlice(3, 10))
}
