package main

import "github.com/01-edu/z01"

func TetraE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for row := 0; row <= y-1; row++ {
		for col := 0; col <= x-1; col++ {
			if row == 0 || row == y-1 {
				if row == 0 && col == 0 {
					z01.PrintRune('A')
				} else if row == 0 && col == x-1 {
					z01.PrintRune('C')
				} else if row == y-1 && col == 0 {
					z01.PrintRune('C')
				} else if row == y-1 && col == x-1 {
					z01.PrintRune('A')
				} else {
					z01.PrintRune('B')
				}
			} else if row > 0 || row < y-1 {
				if col == 0 || col == x-1 {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}

}

func main() {
	TetraE(5, 3)
}
