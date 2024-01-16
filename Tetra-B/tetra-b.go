package main

import "github.com/01-edu/z01"

func TetraB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if row == 0 || row == y-1 {
				if row == 0 && col == 0 {
					z01.PrintRune('/')
				} else if row == 0 && col == x-1 {
					z01.PrintRune('\\')
				} else if row == y-1 && col == 0 {
					z01.PrintRune('\\')
				} else if row == y-1 && col == x-1 {
					z01.PrintRune('/')
				} else {
					z01.PrintRune('*')
				}
			} else if row > 0 || row < y-1 {
				if col == 0 || col == x-1 {
					z01.PrintRune('*')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
	z01.PrintRune('\n')
}

func main() {
	TetraB(5, 3)
}
