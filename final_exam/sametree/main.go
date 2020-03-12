package main

import "fmt"

type TreeNodeL struct {
    Left    *TreeNodeL
    Val     int
    Right   *TreeNodeL
}

func TreeInsert(tree *TreeNodeL, val int) *TreeNodeL {
	if tree == nil {
		return &TreeNodeL{Val: val}
	}
	if tree.Val <= val {
		tree.Right = TreeInsert(tree.Right, val)
	} else {
		tree.Left = TreeInsert(tree.Left, val)
	}
	return tree
}

func IsSameTree(p *TreeNodeL, q *TreeNodeL) bool {
	if p == nil && q == nil {
		return true
	} else if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	// fmt.Println(p.Val, p.Left, p.Left, q.Val, q.Left, q.Right)

	if p.Val != q.Val {
		return false
	}

	left := IsSameTree(p.Left, q.Left)
	right := IsSameTree(p.Right, q.Right)
	if !left || !right {
		return false
	}
	return true
}


func main() {
	t1 := &TreeNodeL{Val: 4}
	TreeInsert(t1, 2)
	TreeInsert(t1, 3)

	t2 := &TreeNodeL{Val: 4}
	TreeInsert(t2, 2)
	TreeInsert(t2, 5)
	
	// fmt.Println(t1.Val, t1.Left, t1.Right)
	// fmt.Println(t2.Val, t2.Left, t2.Right)

	fmt.Println(IsSameTree(t1, t2))
}