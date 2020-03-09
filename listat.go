package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	count := 1
	n := l
	for n != nil {
		if count == pos {
			break
		}
		n = n.Next
		count++
	}
	return n
}
