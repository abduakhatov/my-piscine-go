package main

import (
	"os"
	"github.com/01-edu/z01"
)
// "--cast",
// "github.com/01-edu/z01.PrintRune",
// "os.*",
// "len",
// "--no-lit=[b-mB-M]"


func main() {
	args := os.Args[1:]
	if len(args)  != 1 {
		z01.PrintRune('\n')
		return
	}
	for _, val := range args[0] {
		if !('a' <= val && val <= 'z') || !('A' <= val && val <= 'Z') {
			z01.PrintRune(val)
			continue
		}
		upper := 0
		if ('A' <= val && val <= 'Z') {
			upper = -32
		}

		
		
	}
}