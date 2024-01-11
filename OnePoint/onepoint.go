package main

import "fmt"

func onePoint(p *int) {
	*p = 1
}

func main() {
	x := 4
	onePoint(&x)
	fmt.Println(x)
}
