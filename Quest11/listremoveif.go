package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	for i := l.Head; i != nil; i = i.Next {
		if CompStr(data_ref, i.Data) {
			i = i.Next

		}
	}
}
