package main

import (
	"fmt"
	"github.com/01-edu/z01"
	"os"
)

func PrintHelp() {
	fmt.Printf(
		"--insert" +
			"\n  -i" +
			"\n\tThis flag inserts the string into the string passed as argument." +
			"\n--order" +
			"\n  -o" +
			"\n\tThis flag will behave like a boolean, if it is called it will order the argument.")
}

func PrintRes(text string) {
	str := []rune(text)
	for _, val := range str {
		z01.PrintRune(val)
	}
	z01.PrintRune('\n')
}

func Sort(table []rune, count int) string {
	for i := 0; i < count-1; i++ {
		for j := i + 1; j < count; j++ {
			if table[i] > table[j] {
				tmp := table[i]
				table[i] = table[j]
				table[j] = tmp
			}
		}
	}
	return string(table)
}

func main() {
	args := os.Args[1:]
	toOrder := false
	insert := ""
	text := ""
	help := true
	for _, val := range args {
		if val == "--order" || val == "-o" {
			toOrder = true
		} else if val == "--help" || val == "-h" {
			help = false
			PrintHelp()
		} else if val[:1] == "-" {
			if val[:3] == "-i=" {
				insert = val[3:]
			} else if val[:9] == "--insert=" {
				insert = val[9:]
			}
		} else {
			text += val
		}
	}
	if text == "" && help {
		PrintHelp()
	} else {
		text += insert
		if toOrder {
			counter := 0
			runes := []rune(text)
			for range runes {
				counter++
			}
			text = Sort(runes, counter)
		}
		PrintRes(text)
	}
}
