package tree

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	idxMax := findMax(nums)
	root := &TreeNode{
		Val:   nums[idxMax],
	}
	
	root.Left = constructMaximumBinaryTree(nums[:idxMax])
	root.Right = constructMaximumBinaryTree(nums[idxMax+1:])

	return root
}

func findMax(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	idxMax := 0
	for i := 1; i < len(nums); i++ {
		if nums[idxMax] < nums[i] {
			idxMax = i
		}
	}

	return idxMax
}