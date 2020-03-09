package piscine

func ListLast(l *List) interface{} {
	if l.Tail == nil {
		return nil
	}
	return l.Tail.Data
}
