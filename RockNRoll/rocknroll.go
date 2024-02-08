package main

import "fmt"

func rockNroll(n int) string {
	if n%2 == 0 && n%3 == 0 {
		return "rock and roll"
	}
	if n%3 == 0 {
		return "roll"
	} else if n%2 == 0 {
		return "rock"
	} else if n < 0 {
		return "error: number is negative"
	} else {
		return "error: non divisible"
	}

}

func main() {
	fmt.Println(rockNroll(4))
	fmt.Println(rockNroll(9))
	fmt.Println(rockNroll(6))
	fmt.Println(rockNroll(5))
}
