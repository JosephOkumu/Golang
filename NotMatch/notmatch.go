package main

import "fmt"

func notMatch(a []int) int {
	//create a map to store the count of each element in the array.
	count := make(map[int]int)

	for _, v := range a {
		count[v]++ //Increment the count of the current element in the map.
	}
	for _, v := range a {
		if count[v]%2 != 0 {
			return v
		}
	}
	return -1
}

func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4}
	unmatch := notMatch(a)
	fmt.Println(unmatch)
}
