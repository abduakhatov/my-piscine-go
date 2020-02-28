package main

import (
	piscine ".."
	"github.com/01-edu/z01"
)

func main() {
	piscine.PrintNbrBase(125, "0123456789")
	z01.PrintRune('\n')
	piscine.PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	piscine.PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	piscine.PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	piscine.PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
	piscine.PrintNbrBase(-944556, "CHOUMIisdacat!")
	z01.PrintRune('\n')
	piscine.PrintNbrBase(968670, "WhoAmI?")
	z01.PrintRune('\n')
	piscine.PrintNbrBase(125, "-ab")
	z01.PrintRune('\n')
	z01.PrintRune('>')
	piscine.PrintNbrBase(-9223372036854775808, "")
	z01.PrintRune('\n')
}
