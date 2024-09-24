package src

type binarySearchTree struct {
	root *TreeNode
}

func newBinartSearchTree() *binarySearchTree {
	bst := &binarySearchTree{}
	bst.root = nil
	return bst
}

func (bst *binarySearchTree) getRoot() *TreeNode {
	return bst.root
}

func (bst *binarySearchTree) search(num int) *TreeNode {
	node := bst.root
	for node != nil {
		if node.Val.(int) < num {
			node = node.Right
		} else if node.Val.(int) > num {
			node = node.Left
		} else {
			break
		}
	}
	return node
}

func (bst *binarySearchTree) insert(num int) {
	cur := bst.root
	if cur == nil {
		bst.root = &TreeNode{
			Val:   num,
			Left:  nil,
			Right: nil,
		}
		return
	}
	var pre *TreeNode = nil
	for cur != nil {
		if cur.Val == num {
			return
		}
		pre = cur
		if cur.Val.(int) < num {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	node := &TreeNode{
		Val:   num,
		Left:  nil,
		Right: nil,
	}
	if pre.Val.(int) < num {
		pre.Right = node
	} else {
		pre.Left = node
	}
}

func (bst *binarySearchTree) remove(num int) {
	cur := bst.root
	if cur == nil {
		return
	}
	var pre *TreeNode = nil
	for cur != nil {
		if cur.Val == num {
			return
		}
		pre = cur
		if cur.Val.(int) < num {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	if cur == nil {
		return
	}
	if cur.Left == nil || cur.Right == nil {
		var child *TreeNode = nil
		if cur.Left != nil {
			child = cur.Left
		} else {
			child = cur.Right
		}
		if cur != bst.root {
			if pre.Left == cur {
				pre.Left = child
			} else {
				pre.Right = child
			}
		} else {
			bst.root = child
		}
	} else {
		tmp := cur.Right
		for tmp.Right != nil {
			tmp = tmp.Left
		}
		bst.remove(tmp.Val.(int))
		cur.Val = tmp.Val
	}
}
