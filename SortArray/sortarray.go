package main

import "fmt"

func sortArray(a []string) {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}

	}

}

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	sortArray(result)

	fmt.Println(result)
}
