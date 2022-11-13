package stack

import (
	"algorithm-go/util/maths"
	"fmt"
)

// 接雨水：https://leetcode.cn/problems/trapping-rain-water/
func trap(height []int) int {
	var ans int
	var descStk []int

	for i, v := range height {
		for len(descStk) > 0 && height[descStk[len(descStk)-1]] < v {
			if len(descStk) > 1 {
				ans += (i - descStk[len(descStk)-2] - 1) * (maths.Min(v, height[descStk[len(descStk)-2]]) - height[descStk[len(descStk)-1]])
				fmt.Printf("calc -> top: (%d, %d), next: (%d, %d), ans: %d\n",
					descStk[len(descStk)-1], height[descStk[len(descStk)-1]], i, v, ans)
			}

			fmt.Printf("pop -> top: (%d, %d)\n", descStk[len(descStk)-1], height[descStk[len(descStk)-1]])
			descStk = descStk[:len(descStk)-1]
		}
		descStk = append(descStk, i)
		fmt.Printf("push -> next: (%d, %d)\n", i, v)
	}

	return ans
}
