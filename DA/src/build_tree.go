package src

func dfsBuildTree(preorder []int, inorderMap map[int]int, i, l, r int) *TreeNode {
	if r-l < 0 {
		return nil
	}
	root := newTreeNode(preorder[i])
	m := inorderMap[preorder[i]]
	root.Left = dfsBuildTree(preorder, inorderMap, i+1, l, m-1)
	root.Right = dfsBuildTree(preorder, inorderMap, i+1+m-l, m+1, r)
	return root
}

func buildTree(preorder, inorder []int) *TreeNode {
	inorderMap := make(map[int]int, len(inorder))
	for i := 0; i < len(inorder); i++ {
		inorderMap[preorder[i]] = i
	}
	root := dfsBuildTree(preorder, inorderMap, 0, 0, len(inorder)-1)
	return root
}
