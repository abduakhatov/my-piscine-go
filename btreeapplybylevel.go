package piscine

/*
type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}
*/

type QueueNode struct {
	Data TreeNode
	Next *QueueNode
}

type Queue struct {
	Head *QueueNode
	Tail *QueueNode
}

func QueuePushBack(l *Queue, data *TreeNode) {
	n := &QueueNode{Data: *data}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	l.Tail.Next = n
	l.Tail = l.Tail.Next
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	queue := &Queue{}
	QueuePushBack(queue, root)

	for l := queue.Head; l != nil; l = l.Next {
		f(l.Data.Data)
		if l.Data.Left != nil {
			QueuePushBack(queue, l.Data.Left)
		}
		if l.Data.Right != nil {
			QueuePushBack(queue, l.Data.Right)
		}
	}

}
