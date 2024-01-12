package main

import (
	"fmt"
)

func stringRev(str string) string {
	wrd := ""
	for i := len(str) - 1; i >= 0; i-- {
		ltr := str[i]
		wrd = wrd + string(ltr)

	}
	return wrd

}

func main() {
	str := "Hello World!"
	str = stringRev(str)
	fmt.Println(str)
}
