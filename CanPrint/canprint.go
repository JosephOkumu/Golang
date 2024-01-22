package main

import "fmt"

func canPrint(s string) bool {
	for _, v := range s {
		//Non-printable characters are less than 32 and greater than 126
		if v < 32 || v > 126 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canPrint("Hello"))
	fmt.Println(canPrint("Hello\n"))

}
