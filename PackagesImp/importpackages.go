package main

import (
	"fmt"
	//The t infront of time just says that we are naming the time package t in our program
	t "time"
)

func main() {
	fmt.Println("Learning to import and use packages.")
	fmt.Println(t.Now())
}
