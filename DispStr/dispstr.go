// This program prints one by one the characters of a string on the screen.

package main

import "github.com/01-edu/z01"

func dispStr(str string) {

	for _, v := range str {
		z01.PrintRune(v)
	}
}

func main() {
	dispStr("Hello World!")

}
