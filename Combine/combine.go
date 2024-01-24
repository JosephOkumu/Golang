package main

import "fmt"

func combine(strs []string, sep string) string {
	var result string

	for i, s := range strs {
		if i > 0 {
			result += sep
		}
		result += s
	}

	return result
}

func main() {
	toCombine := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(combine(toCombine, ":"))
}
