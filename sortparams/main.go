package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	args := os.Args[1:]
	for i, a := range args {
		for j, b := range args {
			if a < b {
				args[i], args[j] = args[j], args[i]
			}
		}
	}
	for _, a := range args {
		for _, b := range a {
			z01.PrintRune(b)
			z01.PrintRune('\n')
		}
	}
}
