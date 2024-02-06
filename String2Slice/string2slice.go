package main

import "fmt"

func string2Slice(s string) []int {
	slice := []int{}

	for _, v := range s {
		slice = append(slice, int(v))
	}
	return slice
}

func main() {
	fmt.Println(string2Slice("A quick brown fox jumps over the lazy dog"))
	fmt.Println(string2Slice("Converted this string into an int"))
	fmt.Println(string2Slice("hello THERE"))
}
