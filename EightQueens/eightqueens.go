package main

import "fmt"

func EightQueens(n int, cols []int) {
	if len(cols) == n {
		// Found a solution, print it
		for _, col := range cols {
			fmt.Print(col, "")
		}
		fmt.Println()
		return
	}

	for i := 0; i < n; i++ {
		// Check if placing a queen at i is valid
		valid := true
		for j := 0; j < len(cols); j++ {
			if cols[j] == i ||
				cols[j]-j == i-len(cols) ||
				cols[j]+j == i+len(cols) {
				valid = false
				break
			}
		}

		if valid {
			// Place queen and recur
			cols = append(cols, i)
			EightQueens(n, cols)
			cols = cols[:len(cols)-1]
		}
	}
}

func main() {
	EightQueens(8, []int{})
}
