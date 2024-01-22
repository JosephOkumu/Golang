package main

import "fmt"

func itsaNumber(s string) bool {
	for _, v := range s {
		if v < 48 || v > 57 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(itsaNumber("010203"))
	fmt.Println(itsaNumber("01,02,03"))

}
