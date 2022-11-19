package stack

// 下一个更大元素I：https://leetcode.cn/problems/next-greater-element-i/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))
	descStk := []int{}
	greaters := make(map[int]int, 0)

	for _, v := range nums2 {
		for len(descStk) > 0 && v > descStk[len(descStk)-1] {
			top := descStk[len(descStk)-1]
			descStk = descStk[:len(descStk)-1]
			greaters[top] = v
		}
		descStk = append(descStk, v)
	}

	for i, v := range nums1 {
		if nextG, ok := greaters[v]; ok {
			ans[i] = nextG
		} else {
			ans[i] = -1
		}
	}

	return ans
}

// 下一个更大元素II：https://leetcode.cn/problems/next-greater-element-ii/
func nextGreaterElements(nums []int) []int {
	nums = append(nums, nums...)
	ans := make([]int, len(nums))
	descStk := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		for len(descStk) > 0 && nums[i] > nums[descStk[len(descStk)-1]] {
			ans[descStk[len(descStk)-1]] = nums[i]
			descStk = descStk[:len(descStk)-1]
		}
		descStk = append(descStk, i)
	}

	for i := 0; i < len(descStk); i++ {
		ans[descStk[i]] = -1
	}

	return ans[:len(nums)/2]
}
