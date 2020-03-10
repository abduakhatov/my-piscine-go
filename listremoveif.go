package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	tmp := &List{}
	for i := l.Head; i != nil; i = i.Next {
		if i.Data != data_ref {
			ListPushBack(tmp, i.Data)
		}
	}
	l.Head = tmp.Head
}
