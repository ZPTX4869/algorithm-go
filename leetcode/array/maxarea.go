package array

// 双指针，选出短板，只有短板向内移动可能出现更大的area
func maxArea(height []int) int {
	getSmaller := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	n := len(height)
	if n == 0 {
		return 0
	}

	var res int

	l, r := 0, n-1
	for l < r {
		sHeight := getSmaller(height[l], height[r])
		area := (r - l) * sHeight

		if area > res {
			res = area
		}

		if height[l] == sHeight {
			l++
		} else {
			r--
		}
	}

	return res
}

// 暴力破解+单边优化
//func maxArea(height []int) int {
//	getSmaller := func(x, y int) int {
//		if x < y {
//			return x
//		}
//		return y
//	}
//
//	n := len(height)
//	if n == 0 {
//		return 0
//	}
//
//	var res int
//	var maxLeft int
//
//	for i := 0; i < n-1; i++ {
//		if height[i] <= maxLeft {
//			continue
//		}
//		maxLeft = height[i]
//
//		for j := i + 1; j < n; j++ {
//			area := (j - i) * getSmaller(height[i], height[j])
//			if area > res {
//				res = area
//			}
//		}
//	}
//
//	return res
//}
