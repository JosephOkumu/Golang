// This program finds the number at a particular index in the fibonacci sequence
package main

import "fmt"

func recurseFibo(index int) int {
	if index < 0 {
		return -1
	} else if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	} else {
		//The line below finds the fibo index-1 and fibo of index-2 and sums them.
		return recurseFibo(index-1) + recurseFibo(index-2)
	}
}

func main() {
	arg1 := 6
	fmt.Println(recurseFibo(arg1))

}
