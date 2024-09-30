package tree

import . "practice/pkg"

var res []any

func preOrder(node *TreeNode) {
	if node == nil {
		return
	}
	res = append(res, node.Val)
	preOrder(node.Left)
	preOrder(node.Right)
}

func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	inOrder(node.Left)
	res = append(res, node.Val)
	inOrder(node.Right)
}

func postOrder(node *TreeNode) {
	if node == nil {
		return
	}
	postOrder(node.Left)
	postOrder(node.Right)
	res = append(res, node.Val)
}
