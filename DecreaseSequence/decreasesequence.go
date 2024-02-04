package main

import "github.com/01-edu/z01"

func decSeq() {
	for i := '9'; i >= '0'; i-- {
		for j := '9'; j >= '0'; j-- {
			d := j - 1
			for k := i; k >= '0'; k-- {
				for ; d >= '0'; d-- {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(k)
					z01.PrintRune(d)

					if i > '0' || j > '1' || k > '0' || d > '0' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
				d = '9'
			}
		}
	}
}

func main() {
	decSeq()
}
