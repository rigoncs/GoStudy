package pkg

type TreeNode struct {
	Val    int
	Height int
	Left   *TreeNode
	Right  *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}
