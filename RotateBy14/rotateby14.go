package main

import "github.com/01-edu/z01"

func rotateBy14(str string) string {
	result := ""

	for _, val := range str {

		if val >= 'A' && val < 'Z' {
			val = val + 14

			if val > 'Z' {
				val = (val - 'Z') + ('A' - 1)
			}
		}

		if val >= 'a' && val <= 'z' {
			val += 14
			if val > 'z' {
				val = (val - 'z') + ('a' - 1)
			}
		}
		result += string(val)

	}
	return result
}

func main() {
	result := rotateBy14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
