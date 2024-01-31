package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func iStr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func isInt(n int) {
	val := '0'
	if n/10 > 0 {
		isInt(n / 10)
	}
	for i := 0; i < n%10; i++ {
		val++
	}
	z01.PrintRune(val)

}

func main() {
	points := point{}
	setPoint(&points)
	iStr("x = ")
	isInt(points.x)
	iStr(", y = ")
	isInt(points.y)
	z01.PrintRune('\n')

}
