package piscine

import "fmt"

func ListSize(l *List) int {
	length := 0
	n := l.Head
	for n != nil {
		fmt.Println(n.Data)
		length++
		n = n.Next
	}
	return length
}
