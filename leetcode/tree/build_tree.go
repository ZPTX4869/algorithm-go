package tree

func buildTree(preorder []int, inorder []int) *TreeNode {
	search := func(nums []int, tar int) int {
		for i, v := range nums {
			if v == tar {
				return i
			}
		}

		return -1
	}

	if len(preorder) == 0 {
		return nil
	}

	// 前序数组确定左子树边界元素索引：必须通过中序数组rootVal左侧的元素个数确定
	rootVal := preorder[0]
	idx := search(inorder, rootVal)

	root := &TreeNode{Val: rootVal}
	root.Left = buildTree(preorder[1:1+idx], inorder[:idx])
	root.Right = buildTree(preorder[1+idx:], inorder[idx+1:])

	return root
}
