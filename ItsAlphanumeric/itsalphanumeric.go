// returns true if a string contains alphanumeric or is empty, otherwise it prints false.
package main

import "fmt"

func itsAlphanumeric(str string) bool {
	if len(str) == 0 {
		return true
	}
	for _, v := range str {
		if (v < 65 || v > 90) && (v < 97 || v > 122) && (v < 48 || v > 57) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(itsAlphanumeric("Hello! How are you?"))
	fmt.Println(itsAlphanumeric("HelloHowareyou"))
	fmt.Println(itsAlphanumeric("What's this 4?"))
	fmt.Println(itsAlphanumeric("Whatsthis4"))

}
