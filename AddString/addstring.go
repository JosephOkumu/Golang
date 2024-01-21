//This program concatenates two strings.

package main

import "fmt"

func addString(s1, s2 string) string {
	return s1 + s2
}

func main() {

	fmt.Println(addString("Hello!", " How are you?"))

}
