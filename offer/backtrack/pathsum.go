package backtrack

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var backtrack func(root *TreeNode)
	backtrack = func(root *TreeNode) {
		if root == nil {
			return
		}

		path = append(path, root.Val)
		defer func() {
			path = path[:len(path)-1]
		}()

		if root.Left == nil && root.Right == nil && sumNums(path) == targetSum {
			res = append(res, append([]int{}, path...))
			return
		}

		backtrack(root.Left)
		backtrack(root.Right)
	}

	backtrack(root)

	return res
}

func sumNums(nums []int) (sum int) {
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return
}
