package main

import "fmt"

func simpleJoin(elems []string) string {
	word := ""
	for _, v := range elems {
		word = word + v
	}
	return word
}

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(simpleJoin(elems))
}
