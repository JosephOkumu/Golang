package main

import "fmt"

func sum(s []int, ch chan int) {
	sum := 0

	for _, v := range s {
		sum += v
	}
	ch <- sum // Send the sum to the channel
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0} // Create a slice of integers

	ch := make(chan int) // Create an integer channel

	go sum(s[:len(s)/2], ch) // Start a goroutine to sum the first half of the slice
	go sum(s[len(s)/2:], ch) // Start a goroutine to sum the second half of the slice

	x, y := <-ch, <-ch     // Receive the sums from the channel
	fmt.Println(x, y, x+y) // Print the sums and their total
}
