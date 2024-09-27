package src

type avlTree struct {
	root *TreeNode
}

func newAvlTree() *avlTree {
	return &avlTree{
		root: nil,
	}
}

func (t *avlTree) height(node *TreeNode) int {
	if node != nil {
		return node.Height
	}
	return -1
}

func (t *avlTree) updateHeight(node *TreeNode) {
	lh := t.height(node.Left)
	rh := t.height(node.Right)
	if lh > rh {
		node.Height = lh + 1
	} else {
		node.Height = rh + 1
	}
}

func (t *avlTree) balanceFactor(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return t.height(node.Left) - t.height(node.Right)
}

func (t *avlTree) rightRotate(node *TreeNode) *TreeNode {
	child := node.Left
	grandChild := child.Right
	child.Right = node
	node.Left = grandChild
	t.updateHeight(node)
	t.updateHeight(child)
	return child
}

func (t *avlTree) leftRotate(node *TreeNode) *TreeNode {
	child := node.Right
	grandChild := child.Left
	child.Left = node
	node.Right = grandChild
	t.updateHeight(node)
	t.updateHeight(child)
	return child
}

func (t *avlTree) rotate(node *TreeNode) *TreeNode {
	bf := t.balanceFactor(node)
	if bf > 1 {
		if t.balanceFactor(node.Left) >= 0 {
			return t.rightRotate(node)
		} else {
			node.Left = t.leftRotate(node.Left)
			return t.rightRotate(node)
		}
	}
	if bf < -1 {
		if t.balanceFactor(node.Right) <= 0 {
			return t.leftRotate(node)
		} else {
			node.Right = t.rightRotate(node.Right)
			return t.leftRotate(node)
		}
	}
	return node
}

func (t *avlTree) insert(val int) {
	t.root = t.insertHelper(t.root, val)
}

func (t *avlTree) insertHelper(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return newTreeNode(val)
	}
	if val < node.Val.(int) {
		node.Left = t.insertHelper(node.Left, val)
	} else if val > node.Val.(int) {
		node.Right = t.insertHelper(node.Right, val)
	} else {
		return node
	}
	t.updateHeight(node)
	node = t.rotate(node)
	return node
}

func (t *avlTree) remove(val int) {
	t.root = t.removeHelper(t.root, val)
}

func (t *avlTree) removeHelper(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}
	if val < node.Val.(int) {
		node.Left = t.removeHelper(node.Left, val)
	} else if val > node.Val.(int) {
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
			tmp := node.Right
			for tmp.Left != nil {
				tmp = tmp.Left
			}
			node.Right = t.removeHelper(node.Right, tmp.Val.(int))
			node.Val = tmp.Val
		}
	}
	t.updateHeight(node)
	node = t.rotate(node)
	return node
}

func (t *avlTree) search(val int) *TreeNode {
	cur := t.root
	for cur != nil {
		if cur.Val.(int) < val {
			cur = cur.Right
		} else if cur.Val.(int) > val {
			cur = cur.Left
		} else {
			break
		}
	}
	return cur
}
