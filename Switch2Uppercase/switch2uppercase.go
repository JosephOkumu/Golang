package main

import "fmt"

func switch2Upper(s string) string {
	newstr := ""
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			v = v - 32
		}
		newstr += string(v)
	}
	return newstr

}

func main() {
	fmt.Println(switch2Upper("Hello! How are you?"))
}
