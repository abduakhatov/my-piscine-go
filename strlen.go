package piscine

import "fmt"

func StrLen(str string) int {
	counter := 0
	for ind := range str {
		if str[ind] > 31 {
			counter += 1
		}
	}
	return counter
}

func main() {
	str := "Hello World!"
	nb := StrLen(str)
	fmt.Println(nb)
}