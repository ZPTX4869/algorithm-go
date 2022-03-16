package dynamics

func translateNum(num int) int {
	nums := make([]int, 0)
	for num > 0 {
		curr := num % 10
		num /= 10
		// 每次在头部添加元素，可以得到正序切片
		nums = append([]int{curr}, nums...)
	}

	n := len(nums)
    if n < 2 {
        return 1
    }

	dp := make([]int, n)
	dp[0] = 1
	
	if nums[0] * 10 + nums[1] < 26 {
		dp[1] = 2
	} else {
		dp[1] = 1
	}

	for i := 2; i < n; i++ {
		if nums[i-2] != 0 && nums[i-2] * 10 + nums[i-1] < 26 {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[n-1]
}

func translateNum2(num int) int {
	nums := make([]int, 0)
	for num > 0 {
		curr := num % 10
		num /= 10
		// 每次在头部添加元素，可以得到正序切片
		nums = append([]int{curr}, nums...)
	}

	n := len(nums)
	if n == 0 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		if nums[i-2] != 0 && nums[i-2] * 10 + nums[i-1] < 26 {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[n]
}