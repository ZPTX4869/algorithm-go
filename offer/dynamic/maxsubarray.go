package dynamic

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	getMax := func(x, y int) int {
		if x > y {
			return x
		}

		return y
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	max := dp[0]

	for i := 1; i < n; i++ {
		dp[i] = getMax(dp[i-1]+nums[i], nums[i])

		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}

func maxSubArray2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	getMax := func(x, y, z int) int {
		max := x
		if y > max {
			max = y
		}
		if z > max {
			max = z
		}

		return max
	}

	var helper func(i, j int) int
	helper = func(left, right int) int {
		if left == right {
			return nums[left]
		}

		mid := (right-left)/2 + left

		return getMax(helper(left, mid), helper(mid+1, right), maxCrossArr(nums, left, mid, right))
	}

	return helper(0, len(nums)-1)
}

func maxCrossArr(nums []int, left, mid, right int) int {
	leftMax, leftSum := nums[mid], nums[mid]
	rightMax, rightSum := nums[mid+1], nums[mid+1]

	for i := mid - 1; i >= left; i-- {
		leftSum += nums[i]
		if leftSum > leftMax {
			leftMax = leftSum
		}
	}

	for i := mid + 2; i <= right ; i++ {
		rightSum += nums[i]
		if rightSum > rightMax {
			rightMax = rightSum
		}
	}

	return leftMax + rightMax
}
