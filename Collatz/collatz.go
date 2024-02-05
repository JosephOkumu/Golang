package main

import "fmt"

func collatzTally(first int) int {
	stps := 0
	if first <= 0 {
		return -1
	}

	for first > 1 {
		if first%2 == 0 {
			first = first / 2
		} else if first%2 != 0 {
			first = 3*first + 1
		}
		stps++
	}
	return stps

}

func main() {
	stps := collatzTally(12)
	fmt.Println(stps)
}
