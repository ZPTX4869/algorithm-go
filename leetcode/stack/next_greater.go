package stack

// 下一个更大元素I：https://leetcode.cn/problems/next-greater-element-i/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	descStk := []int{}
	greaters := make(map[int]int, 0)

	for _, v := range nums2 {
		for len(descStk) > 0 && v > descStk[len(descStk)-1] {
			top := descStk[len(descStk)-1]
			descStk = descStk[:len(descStk)-1]
			greaters[top] = v
			continue
		}
		descStk = append(descStk, v)
	}

	ans := make([]int, len(nums1))
	for i, v := range nums1 {
		if nextG, ok := greaters[v]; ok {
			ans[i] = nextG
		} else {
			ans[i] = -1
		}
	}

	return ans
}
