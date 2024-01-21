//Checks if string passed contains only lowercase letters

package main

import "fmt"

func itsLower(str string) bool {
	for _, v := range str {
		if v < 97 || v > 122 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(itsLower("bumper"))
	fmt.Println(itsLower("bumper!"))

}
