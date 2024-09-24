package src

import "container/list"

type TreeNode struct {
	Val   any
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []any {
	queue := *list.New()
	queue.PushBack(root)
	nums := make([]any, 0)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		nums = append(nums, node.Val)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return nums
}
