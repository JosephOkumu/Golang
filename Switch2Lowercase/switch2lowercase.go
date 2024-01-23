package main

import "fmt"

func switch2Lower(str string) string {
	newstr := ""
	for _, v := range str {
		if v >= 65 && v <= 90 {
			v = v + 32
		}
		newstr += string(v)
	}
	return newstr
}

func main() {
	fmt.Println(switch2Lower("What the Heck! What is Rooftop?"))

}
