package main

import "fmt"

func rmWhitespaces(str string) []string {
	var check []string
	var new string

	for _, v := range str {
		if v != 32 && v != '\t' && v != '\n' && v != '\r' {
			new = new + string(v)
		} else if new != "" {
			check = append(check, new)
			new = ""
		}
	}
	if new != "" {
		check = append(check, new)
	}
	return check
}

func main() {
	fmt.Printf("%#v\n", rmWhitespaces("Who are you going to be in the near future"))
}
