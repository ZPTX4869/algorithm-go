package trees

import "math"

// Note: this is a binary search tree
func verifyPostorder(postorder []int) bool {
	var helper func(postorder []int, lower, upper int) bool
	helper = func(postorder []int, lower, upper int) bool {
		n := len(postorder)
		if n < 1 {
			return true
		}

		if postorder[n-1] < lower || postorder[n-1] > upper {
			return false
		}

		right := n - 1
		for i, v := range postorder {
			if v > postorder[n-1] {
				right = i
				break
			}
		}

		return helper(postorder[:right], lower, postorder[n-1]) && 
		helper(postorder[right:n-1], postorder[n-1], upper)
	}

	n := len(postorder)
	if n == 0 {
		return true
	}

	return helper(postorder, math.MinInt, math.MaxInt)
}
