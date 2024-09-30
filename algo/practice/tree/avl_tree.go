package tree

import (
	. "practice/pkg"
)

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

func (t *aVLTree) balanceFactor(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return t.height(node.Left) - t.height(node.Right)
}

func (t *aVLTree) rightRotate(node *TreeNode) *TreeNode {
	child := node.Left
	grandChild := child.Right
	child.Right = node
	node.Left = grandChild
	t.updateHeight(node)
	t.updateHeight(child)
	return child
}

func (t *aVLTree) leftRotate(node *TreeNode) *TreeNode {
	child := node.Right
	grandChild := child.Left
	child.Left = node
	node.Right = grandChild
	t.updateHeight(node)
	t.updateHeight(child)
	return child
}

func (t *aVLTree) rotate(node *TreeNode) *TreeNode {
	bf := t.balanceFactor(node)
	if bf > 1 {
		if t.balanceFactor(node.Left) >= 0 {
			node = t.rightRotate(node)
		} else {
			node.Left = t.leftRotate(node.Left)
			node = t.rightRotate(node)
		}
	} else if bf < -1 {
		if t.balanceFactor(node.Right) <= 0 {
			node = t.leftRotate(node)
		} else {
			node.Right = t.rightRotate(node.Right)
			node = t.leftRotate(node)
		}
	}
	return node
}

func (t *aVLTree) insert(val int) {
	t.root = t.insertHelper(t.root, val)
}

func (t *aVLTree) insertHelper(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return NewTreeNode(val)
	}
	if val < node.Val {
		node.Left = t.insertHelper(node.Left, val)
	} else if val > node.Val {
		node.Right = t.insertHelper(node.Right, val)
	} else {
		return node
	}
	t.updateHeight(node)
	node = t.rotate(node)
	return node
}

func (t *aVLTree) remove(val int) {
	t.root = t.removeHelper(t.root, val)
}

func (t *aVLTree) removeHelper(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}
	if val < node.Val {
		node.Left = t.removeHelper(node.Left, val)
	} else if val > node.Val {
		node.Right = t.removeHelper(node.Right, val)
	} else {
		if node.Left == nil || node.Right == nil {
			child := node.Left
			if node.Right != nil {
				child = node.Right
			}
			if child == nil {
				return nil
			} else {
				node = child
			}
		} else {
			temp := node.Right
			for temp != nil {
				temp = temp.Left
			}
			node.Right = t.removeHelper(node.Right, temp.Val)
			node.Val = temp.Val
		}
	}
	t.updateHeight(node)
	node = t.rotate(node)
	return node
}

func (t *aVLTree) search(val int) *TreeNode {
	curr := t.root
	for curr != nil {
		if val < curr.Val {
			curr = curr.Left
		} else if val > curr.Val {
			curr = curr.Right
		} else {
			break
		}
	}
	return curr
}
