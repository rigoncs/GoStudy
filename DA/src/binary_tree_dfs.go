package src

var nums []any

func preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	nums = append(nums, root.Val)
	preOrder(root.Left)
	preOrder(root.Right)
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left)
	nums = append(nums, root.Val)
	inOrder(root.Right)
}

func postOrder(root *TreeNode) {
	if root == nil {
		return
	}
	postOrder(root.Left)
	postOrder(root.Right)
	nums = append(nums, root.Val)
}
