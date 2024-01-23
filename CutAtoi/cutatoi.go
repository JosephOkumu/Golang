package main

import "fmt"

func cutAtoi(str string) int {
	numb := 0
	isNegative := false

	for _, c := range str {
		//check for numbers
		if c >= '0' && c <= '9' {
			digit := int(c - '0')
			numb = numb*10 + digit
		} else if c == '-' && numb == 0 {
			isNegative = true
		}
	}
	if isNegative {
		numb = -numb
	}
	return numb
}

func main() {
	fmt.Println(cutAtoi("12345"))
	fmt.Println(cutAtoi("str123ing45"))
	fmt.Println(cutAtoi("012 345"))
	fmt.Println(cutAtoi("Hello World!"))
	fmt.Println(cutAtoi("sd+x1fa2W3s4"))
	fmt.Println(cutAtoi("sd-x1fa2W3s4"))
	fmt.Println(cutAtoi("sdx1-fa2W3s4"))
	fmt.Println(cutAtoi("sdx1+fa2W3s4"))

}
