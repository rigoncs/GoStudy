package src

func preOrderI(root *TreeNode, res *[]*TreeNode) {
	if root == nil {
		return
	}
	if (root.Val).(int) == 7 {
		*res = append(*res, root)
	}
	preOrderI(root.Left, res)
	preOrderI(root.Right, res)
}

func preOrderII(root *TreeNode, res *[][]*TreeNode, path *[]*TreeNode) {
	if root == nil {
		return
	}
	*path = append(*path, root)
	if root.Val.(int) == 7 {
		*res = append(*res, append([]*TreeNode{}, *path...))
	}
	preOrderII(root.Left, res, path)
	preOrderII(root.Right, res, path)
	*path = (*path)[:len(*path)-1]
}

func preOrderIII(root *TreeNode, res *[][]*TreeNode, path *[]*TreeNode) {
	if root == nil || root.Val == 3 {
		return
	}
	*path = append(*path, root)
	if root.Val.(int) == 7 {
		*res = append(*res, append([]*TreeNode{}, *path...))
	}
	preOrderII(root.Left, res, path)
	preOrderII(root.Right, res, path)
	*path = (*path)[:len(*path)-1]
}

// template
func isSolution(state *[]*TreeNode) bool {
	return len(*state) != 0 && (*state)[len(*state)-1].Val == 7
}

func recordSolution(state *[]*TreeNode, res *[][]*TreeNode) {
	*res = append(*res, append([]*TreeNode{}, *state...))
}

func isValid(state *[]*TreeNode, choice *TreeNode) bool {
	return choice != nil && choice.Val != 3
}

func makeChoice(state *[]*TreeNode, choice *TreeNode) {
	*state = append(*state, choice)
}

func undoChoice(state *[]*TreeNode, choice *TreeNode) {
	*state = (*state)[:len(*state)-1]
}

func backTrack(state *[]*TreeNode, choices *[]*TreeNode, res *[][]*TreeNode) {
	if isSolution(state) {
		recordSolution(state, res)
	}
	for _, choice := range *choices {
		if isValid(state, choice) {
			makeChoice(state, choice)
			temp := make([]*TreeNode, 0)
			temp = append(temp, choice.Left, choice.Right)
			backTrack(state, &temp, res)
			undoChoice(state, choice)
		}
	}
}
