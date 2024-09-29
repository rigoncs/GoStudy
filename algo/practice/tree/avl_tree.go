package tree

import . "practice/pkg"

type aVLTree struct {
	root *TreeNode
}

func newAVLTree() *aVLTree {
	return &aVLTree{root: nil}
}

func (t *aVLTree) height(node *TreeNode) int {
	if node != nil {
		return node.Height
	}
	return -1
}

func (t *aVLTree) updateHeight(node *TreeNode) {
	leftHeight := t.height(node.Left)
	rightHeight := t.height(node.Right)
	if leftHeight > rightHeight {
		node.Height = leftHeight + 1
	} else {
		node.Height = rightHeight + 1
	}
}
