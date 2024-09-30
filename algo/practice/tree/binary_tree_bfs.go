package tree

import (
	"container/list"
	. "practice/pkg"
)

func levelOrder(root *TreeNode) []any {
	queue := list.New()
	queue.PushBack(root)
	res := make([]any, 0)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		res = append(res, node.Val)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return res
}
