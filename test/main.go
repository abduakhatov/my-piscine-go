package main

// import piscine ".."

import "github.com/01-edu/z01"
import "fmt"

func main() {
	// Task 4
	// piscine.IsNegative(1)
	// piscine.IsNegative(0)
	// piscine.IsNegative(-1)

	// Task 5
	// piscine.PrintComb()
	
	// Task 6
	// piscine.PrintComb2()
	c := 'A' // rune (characters in Go are represented using `rune` data type)
	asciiValue := int(c)

	fmt.Printf("Ascii Value of %c = %d\n", c, asciiValue)
	z01.PrintRune(rune(asciiValue))
}
