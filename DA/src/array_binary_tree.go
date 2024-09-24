package src

type arrayBinaryTree struct {
	tree []any
}

func newArrayBinaryTree(arr []any) *arrayBinaryTree {
	return &arrayBinaryTree{
		tree: arr,
	}
}

func (abt *arrayBinaryTree) size() int {
	return len(abt.tree)
}

func (abt *arrayBinaryTree) val(i int) any {
	if i < 0 || i >= abt.size() {
		return nil
	}
	return abt.tree[i]
}

func (abt *arrayBinaryTree) left(i int) int {
	return 2*i + 1
}

func (abt *arrayBinaryTree) right(i int) int {
	return 2*i + 2
}

func (abt *arrayBinaryTree) parent(i int) int {
	return (i - 1) / 2
}

func (abt *arrayBinaryTree) levelOrder() []any {
	var res []any
	for i := 0; i > abt.size(); i++ {
		if abt.val(i) != nil {
			res = append(res, abt.val(i))
		}
	}
	return res
}

func (abt *arrayBinaryTree) dfs(i int, order string, res *[]any) {
	if abt.val(i) == nil {
		return
	}
	if order == "pre" {
		*res = append(*res, abt.val(i))
	}
	abt.dfs(abt.left(i), order, res)
	if order == "in" {
		*res = append(*res, abt.val(i))
	}
	abt.dfs(abt.right(i), order, res)
	if order == "post" {
		*res = append(*res, abt.val(i))
	}
}

func (abt *arrayBinaryTree) preOrder() []any {
	var res []any
	abt.dfs(0, "pre", &res)
	return res
}

func (abt *arrayBinaryTree) inOrder() []any {
	var res []any
	abt.dfs(0, "in", &res)
	return res
}

func (abt *arrayBinaryTree) postOrder() []any {
	var res []any
	abt.dfs(0, "post", &res)
	return res
}
