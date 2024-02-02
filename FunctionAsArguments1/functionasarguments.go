package main

import "fmt"

func Map(f func(int) bool, a []int) []bool {
	res := make([]bool, len(a))

	for i, v := range a {
		res[i] = f(v)
	}
	return res
}

func itsPrime(n int) bool {
	if n == 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := Map(itsPrime, a)
	fmt.Println(result)

}
