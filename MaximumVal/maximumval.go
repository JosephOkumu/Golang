//This program displays the maximum value found in a slice

package main

import "fmt"

func maxi(a []int) int {
	if a == nil {
		return 0
	}

	max := a[0]
	for _, val := range a {
		if val > max {
			max = val
		}
	}
	return max
}

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	fmt.Println(maxi(a))
}
