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
	points := &point{}

	setPoint(points)
	str := "x = " + string(points.x) + ", y = " + string(points.y) + "\n"

	for _, val := range str {
		z01.PrintRune(val)
	}
}
