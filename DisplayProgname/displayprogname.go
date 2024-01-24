package main

import (
	"fmt"
	"os"
)

func displayProgname() {
	args := os.Args[0]
	fmt.Println(args)
}

func main() {
	displayProgname()
}
