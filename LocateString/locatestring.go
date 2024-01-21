package main

import "fmt"

func locateString(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}
	if len(s) == 0 {
		return -1
	}
	for i := 0; i < len(s); i++ {
		if s[i] == toFind[0] {
			if len(toFind) == 1 {
				return i
			}
			if len(s[i:]) >= len(toFind) {
				if s[i:i+len(toFind)] == toFind {
					return i
				}
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(locateString("Hello!", "l"))
	fmt.Println(locateString("Salut!", "alu"))
	fmt.Println(locateString("Ola!", "hOl"))

}
