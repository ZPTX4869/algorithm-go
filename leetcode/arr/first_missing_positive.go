package arr

// 缺失的第一个整数：https://leetcode.cn/problems/first-missing-positive/submissions/
func firstMissingPositive(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		for nums[i] >= 1 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			temp := nums[nums[i]-1]
			nums[nums[i]-1] = nums[i]
			nums[i] = temp
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}
