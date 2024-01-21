// prints true if string contains uppercase letters
package main

import "fmt"

func itsUpper(s string) bool {
	for _, v := range s {
		if v < 65 || v >= 90 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(itsUpper("HELLO"))
	fmt.Println(itsUpper("HELLO!"))

}
