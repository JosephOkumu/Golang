/*This program prints a combination of two different 2 digit nummbers in ascending order
separated by a comma and a space */

package main

import "github.com/01-edu/z01"

func main() {
	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '9'; b++ {
			for i := '0'; i <= '9'; i++ {
				for j := '0'; j <= '9'; j++ {
					if a < i || (a == i && b < j) {
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(' ')
						z01.PrintRune(i)
						z01.PrintRune(j)

						if a != '9' || b != '8' || i != '9' || j != '9' {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}

	}
	z01.PrintRune('\n')
}
