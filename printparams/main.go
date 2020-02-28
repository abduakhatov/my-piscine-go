package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	name := os.Args
	for _, a := range name[1:] {
		for _, b := range a {
			z01.PrintRune(b)
		}
		z01.PrintRune('\n')
	}
}
