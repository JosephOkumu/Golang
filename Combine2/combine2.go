// combining strings of a slice using the sep argument
package main

import "fmt"

func combine2(strs []string, sep string) string {
	wrd := ""

	for i, v := range strs {
		if i < len(strs)-1 {
			wrd = wrd + v + sep

		} else {
			wrd = wrd + v
		}
	}
	return wrd
}

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(combine2(toConcat, ":"))
}
