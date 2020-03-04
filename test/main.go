package main

import (
	hackathon "hackathon"

	"github.com/01-edu/z01"
)

func main() {
	result := hackathon.Rot14("Hello How are You")
	arrayRune := []rune(result)

	for _, s := range arrayRune {
		z01.PrintRune(s)
	}
	z01.PrintRune('\n')
}
