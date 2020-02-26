package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	name := os.Args[0]
	for _, val := range name {
		z01.PrintRune(val)
	}
	z01.PrintRune('\n')
}
