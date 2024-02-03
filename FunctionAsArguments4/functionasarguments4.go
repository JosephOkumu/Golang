package main

import "fmt"

func isSorted(f func(a, b int) int, a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) == 1 {
			return false
		}
	}
	return true
}

func ifNum(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 1
	}
}

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	result1 := isSorted(ifNum, a1)
	result2 := isSorted(ifNum, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
