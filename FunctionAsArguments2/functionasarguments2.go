//This program checks if a slice contains an int and returns true if so, else it returns false.

package main

import (
	"fmt"
	"strconv"
)

func Check(f func(string) bool, a []string) bool {

	for _, v := range a {
		if f(v) {
			return true
		}
	}
	return false
}

func itsNum(str string) bool {
	i, _ := strconv.Atoi(str)

	if i > 0 && i < 9 {
		return true
	}
	return false
}

func main() {
	a1 := []string{"Hello", "how", "are", "you"}
	a2 := []string{"This", "is", "4", "you"}

	result1 := Check(itsNum, a1)
	result2 := Check(itsNum, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
