package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput() []rune {
	reader := bufio.NewReader(os.Stdin)
	var output []rune
	if stat, _ := os.Stdin.Stat(); (stat.Mode() & os.ModeCharDevice) > 0 {
		fmt.Println("2")
		return output
	}
	for {
		if consoleOutput, _, err := reader.ReadRune(); err != nil {
			break
		} else {
			output = append(output, consoleOutput)
		}
	}
	return output
}

func Validate(output []rune) bool {
	if output[0] != 'o' && output[0] != '/' && output[0] != 'A' {
		return false
	}
	return true
}

func GetParams(output []rune) (x, y, xEnd, yEnd int) {
	x, y = 0, 0
	for _, r := range output {
		if r == '\n' {
			break
		}
		x++
	}
	for _, r := range output {
		if r == '\n' {
			y++
		}
	}
	xEnd = x*y + y - x - 1
	yEnd = x*y + y - 2
	return x, y, xEnd, yEnd
}

func printInput(output []rune) {
	x, y, xEnd, yEnd := GetParams(output)
	switch {
	case output[0] == 'o':
		fmt.Printf("[raid1a] [%d] [%d]", x, y)
		break
	case output[0] == '/':
		fmt.Printf("[raid1b] [%d] [%d]", x, y)
		break
	case output[x-1] == 'A' && output[y-1] == 'A' && x == 1 && y == 1:
		fmt.Printf("[raid1c] [%d] [%d] || [raid1d] [%d] [%d] || [raid1e] [%d] [%d]", x, y, x, y, x, y)
		break
	case output[0] == 'A' && output[xEnd] == 'C' && x == 1 && y > 0:
		fmt.Printf("[raid1c] [%d] [%d] || [raid1e] [%d] [%d]", x, y, x, y)
		break
	case output[0] == 'A' && output[yEnd] == 'C' && y == 1 && x > 0:
		fmt.Printf("[raid1d] [%d] [%d] || [raid1e] [%d] [%d]", x, y, x, y)
		break
	case output[0] == 'A' && output[x-1] == 'A':
		fmt.Printf("[raid1c] [%d] [%d]", x, y)
		break
	case output[0] == 'A' && output[xEnd] == 'A' && output[yEnd] == 'C':
		fmt.Printf("[raid1d] [%d] [%d]", x, y)
		break
	case output[0] == 'A' && output[xEnd] == 'C' && output[yEnd] == 'A':
		fmt.Printf("[raid1e] [%d] [%d]", x, y)
		break
	default:
		fmt.Printf("Not a Raid function")
	}
	fmt.Println()
}

func main() {
	var output []rune = ReadInput()
	if len(output) > 0 {
		if !Validate(output) {
			// if len(output) > 0 && output[0] == 10 {
			//     fmt.Println("Not a Raid function")
			// } else {
			fmt.Println("Not a Raid function")
			// }
			return
		}
		printInput(output)
	} else {
		fmt.Println("Not a Raid function")
	}
}
