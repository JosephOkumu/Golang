package main

import (
	"fmt"
	"strconv"
)

func tallyIf(f func(string) bool, tab []string) int {
	tally := 0
	for _, v := range tab {

		if itsNum(v) {
			tally++
		}
	}
	return tally
}

func itsNum(str string) bool {
	i, _ := strconv.Atoi(str)
	if i > 0 && i < 9 {
		return true
	}
	return false

}

func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This", "1", "is", "4", "you"}
	answer1 := tallyIf(itsNum, tab1)
	answer2 := tallyIf(itsNum, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}
