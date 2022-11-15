package stack

import "fmt"

// 柱状图中最大的矩形：https://leetcode.cn/problems/largest-rectangle-in-histogram/submissions/
func largestRectangleArea(heights []int) int {
	n := len(heights)
	lb, rb := make([]int, n), make([]int, n) // left bound and right bound
	lstk, rstk := []int{-1}, []int{n}        // two stks from both side to store index of elem

	for i := 0; i <= n-1; i++ {
		for len(lstk) > 1 && heights[i] < heights[lstk[len(lstk)-1]] {
			lstk = lstk[:len(lstk)-1]
		}
		lb[i] = lstk[len(lstk)-1]
		lstk = append(lstk, i)
	}

	for i := n - 1; i >= 0; i-- {
		for len(rstk) > 1 && heights[i] <= heights[rstk[len(rstk)-1]] {
			rstk = rstk[:len(rstk)-1]
		}
		rb[i] = rstk[len(rstk)-1]
		rstk = append(rstk, i)
	}

	fmt.Println(lb)
	fmt.Println(rb)

	var ans int
	for i := 0; i <= n-1; i++ {
		area := heights[i] * (rb[i] - lb[i] - 1)
		if area > ans {
			ans = area
		}
	}

	return ans
}
