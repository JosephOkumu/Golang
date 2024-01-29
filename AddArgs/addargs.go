package main

import "fmt"

func addArgs(args []string) string {
	result := ""

	for i := 0; i < len(args); i++ {
		result += args[i] + "\n"
	}
	return result
}

func main() {
	test := []string{"Here", "comes", "your", "destiny"}
	fmt.Println(addArgs(test))
}
