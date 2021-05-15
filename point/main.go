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

func main() {
	xStr := "x = "
	yStr := "y = "
	points := &point{}

	setPoint(points)

	xmassiv := []rune{}
	ymassiv := []rune{}

	xVal := points.x
	yVal := points.y

	for xVal != 0 {
		xmassiv = append(xmassiv, rune(xVal%10))
		xVal /= 10
	}
	for _, val := range xStr {
		z01.PrintRune(rune(val))
	}
	for i := len(xmassiv) - 1; i >= 0; i-- {
		z01.PrintRune(xmassiv[i] + 48)
	}
	z01.PrintRune(',')
	z01.PrintRune(' ')

	for yVal != 0 {
		ymassiv = append(ymassiv, rune(yVal%10))
		yVal /= 10
	}
	for _, val := range yStr {
		z01.PrintRune(rune(val))
	}
	for i := len(ymassiv) - 1; i >= 0; i-- {
		z01.PrintRune(ymassiv[i] + 48)
	}
	z01.PrintRune('\n')
}
