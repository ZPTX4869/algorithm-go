package trees

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}
	
	rootIdx := 0
	for idx, num := range inorder {
		if num == preorder[0] {
			rootIdx = idx
			break
		}
	}

	root.Left = buildTree(preorder[1:len(inorder[:rootIdx])+1], inorder[:rootIdx])
	root.Right = buildTree(preorder[len(inorder[:rootIdx])+1:], inorder[rootIdx+1:])

	return root
}
