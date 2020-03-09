package piscine

func ListSize(l *List) int {
	length := 0
	n := l.Head
	for n != nil {
		length++
		n = n.Next
	}
	return length
}
