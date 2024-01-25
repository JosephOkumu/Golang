package main

import "fmt"

func main() {
	for count := 0; count < 20; count++ {
		if count == 8 {
			// moves/skips to the next count ==9 when count count == 8.
			continue
		}
		if count == 15 {
			//Ends the iteration.
			break
		}
		fmt.Println(count)
	}

}
