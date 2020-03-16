package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BtreeAtoi(nbr string) int {
	var res int = 0
	var sign int = 1
	if nbr != "" && nbr[0] == '-' {
		nbr = nbr[1:]
		sign *= -1
	} else if nbr != "" && nbr[0] == '+' {
		nbr = nbr[1:]
	}
	for _, digit := range nbr {
		tmp := 0
		for i := '1'; i <= digit; i++ {
			tmp++
		}
		res = res*10 + (tmp * sign)
	}
	return res
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	// fmt.Println(root.Data, data, BtreeAtoi(root.Data) > BtreeAtoi(data), BtreeAtoi(root.Data) <= BtreeAtoi(data))
	if root.Data <= data {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	} else {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	}
	return root
}
