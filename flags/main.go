package main

import (
	"os"
	"fmt"
	// "github.com/01-edu/z01"
)

func PrintHelp() {
	fmt.Printf(
		"--insert" + 
		"\n  -i" +
		"\n    This flag inserts the string into the string passed as argument."+
		"\n--order"+
		"\n  -o"+
		"\n    This flag will behave like a boolean, if it is called it will order the argument.")
}

func PrintFlags() {
	return
}

func main() {
	args := os.Args[1:]
	fmt.Println(args)

	a := args[0] == []string{}

	if args[0] == "--help" || args[0] == "-h" {
		PrintHelp()
	} else {
		PrintFlags()
	}
}
