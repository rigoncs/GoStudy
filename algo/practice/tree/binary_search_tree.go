package tree

import . "practice/pkg"

type binarySearchTree struct {
	root *TreeNode
}

func newBinarySearchTree() *binarySearchTree {
	return &binarySearchTree{
		root: nil,
	}
}

func (bst *binarySearchTree) getRoot() *TreeNode {
	return bst.root
}

func (bst *binarySearchTree) search(num int) *TreeNode {
	curr := bst.getRoot()
	for curr != nil {
		if num < curr.Val {
			curr = curr.Left
		} else if num > curr.Val {
			curr = curr.Right
		} else {
			break
		}
	}
	return curr
}

func (bst *binarySearchTree) insert(num int) {
	curr := bst.getRoot()
	if curr == nil {
		bst.root = NewTreeNode(num)
		return
	}
	var pre *TreeNode
	for curr != nil {
		if num == curr.Val {
			return
		} else if num < curr.Val {
			pre = curr
			curr = curr.Left
		} else {
			pre = curr
			curr = curr.Right
		}
	}
	node := NewTreeNode(num)
	if pre.Val > num {
		pre.Left = node
	} else {
		pre.Right = node
	}
}

func (bst *binarySearchTree) remove(num int) {
	curr := bst.getRoot()
	if curr == nil {
		return
	}
	var pre *TreeNode
	for curr != nil {
		if num == curr.Val {
			break
		}
		pre = curr
		if num < curr.Val {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
	if curr == nil {
		return
	}
	if curr.Left == nil || curr.Right == nil {
		var child *TreeNode
		if curr.Left != nil {
			child = curr.Left
		} else {
			child = curr.Right
		}
		if curr == bst.root {
			bst.root = child
		} else {
			if pre.Left == curr {
				pre.Left = child
			} else {
				pre.Right = child
			}
		}
	} else {
		temp := curr.Right
		for temp != nil {
			temp = temp.Left
		}
		bst.remove(temp.Val)
		curr.Val = temp.Val
	}
}
