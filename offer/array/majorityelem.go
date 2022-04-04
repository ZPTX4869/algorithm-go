package array

// 分治法
func majorityElement(nums []int) int {
	var divide func(low, high int) int
	var conquer func(low, high, leftRes, rightRes int) int

	divide = func(low, high int) int {
		if low == high {
			return nums[low]
		}

		mid := (high-low)/2 + low
		leftRes := divide(low, mid)
		rightRes := divide(mid+1, high)

		return conquer(low, high, leftRes, rightRes)
	}

	conquer = func(low, high, leftRes, rightRes int) int {
		if leftRes == rightRes {
			return leftRes
		}

		leftCnt, rightCnt := 0, 0
		for i := low; i <= high; i++ {
			if nums[i] == leftRes {
				leftCnt += 1
			} else if nums[i] == rightRes {
				rightCnt += 1
			}
		}

		if leftCnt > rightCnt {
			return leftRes
		} else {
			return rightRes
		}
	}

	return divide(0, len(nums)-1)
}

// 摩尔计数法
func majorityElement2(nums []int) int {
	cand := nums[0]
	cnt := 0

	for _, num := range nums {
		if cand == num {
			cnt += 1
		} else {
			cnt -= 1
		}

		// 更换选举人
		if cnt == 0 {
			cand = num
			cnt = 1
		}
	}

	return cand
}
