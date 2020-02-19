package piscine

import (
	"github.com/01-edu/z01"
	"fmt"
)

func PrintNbr(n int) {
	t := fmt.Sprintf("%d", n)
	// fmt.Printf(t) # also prints
	slice := []rune(t)
	for i := 0; i < len(slice); i++ {
		z01.PrintRune(slice[i])
	}
	z01.PrintRune('\n')
}
