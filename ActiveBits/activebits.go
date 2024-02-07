package main

import (
	"fmt"
	"strconv"
)

func activeBits(n int) int {
	bits := strconv.FormatInt(7, 2)
	//fmt.Println(bits)
	count := 0
	for _, v := range bits {
		if v > 0 {
			count++
		}

	}
	return count
}

func main() {
	fmt.Println(activeBits(7))
}
