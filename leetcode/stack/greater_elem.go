package stack

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	greaters := make(map[int]int, len(nums2))

	for i := range nums2 {
		greaters[nums2[i]] = -1
	}

	var stk []int
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stk) > 0 && stk[len(stk)-1] <= nums2[i] {
			stk = stk[:len(stk)-1]
		}

		if len(stk) > 0 {
			greaters[nums2[i]] = stk[len(stk)-1]
		}

		stk = append(stk, nums2[i])
	}

	for i, v := range nums1 {
		res[i] = greaters[v]
	}

	return res
}
