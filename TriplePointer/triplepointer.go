package main

import "fmt"

func triplePoint(p ***int) {
	***p = 5
}

func main() {
	a := 0
	b := &a
	c := &b

	triplePoint(&c)
	fmt.Println(a)
}
