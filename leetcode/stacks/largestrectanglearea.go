package stacks

// 单调栈求解，时间复杂度为 O(N)
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	heights = append([]int{0}, heights...)

	max := 0
	stack := make([]int, 0)
	for i := 0; i <= len(heights); i++ {
		var cur int
		if i == len(heights) {
			cur = 0
		} else {
			cur = heights[i]
		}

		for len(stack) != 0 && heights[stack[len(stack)-1]] > cur {
			head := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			height := heights[head]

			if len(stack) != 0 {
				peek := stack[len(stack)-1]
				max = Max(max, (i-peek-1)*height)
			}
		}

		stack = append(stack, i)
	}

	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 暴力求解，会超时
//func largestRectangleArea(heights []int) int {
//	if len(heights) == 0 {
//		return 0
//	}
//
//	maxArea := 0
//	for i := 0; i < len(heights); i++ {
//		if currArea := calcRectangleArea(heights, i); currArea > maxArea {
//			maxArea = currArea
//		}
//	}
//
//	return maxArea
//}
//
//func calcRectangleArea(heights []int, root int) int {
//	areaSum := heights[root]
//
//	left := root
//	for left-1 >= 0 && heights[left-1] >= heights[root] {
//		areaSum += heights[root]
//		left = left - 1
//	}
//
//	right := root
//	for right+1 < len(heights) && heights[right+1] >= heights[root] {
//		areaSum += heights[root]
//		right = right + 1
//	}
//
//	return areaSum
//}
