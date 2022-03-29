package backtracks

import (
	"sort"
)

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	paths := make([]int, 0)

	var backtrack func(start int)
	backtrack = func(start int) {
		res = append(res, append([]int{}, paths...))

		for i := start; i < len(nums); i++ {
			paths = append(paths, nums[i])
			backtrack(i + 1)
			paths = paths[:len(paths)-1]
		}
	}

	backtrack(0)
	return res
}

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	paths := make([]int, 0)

	var backtrack func(start int)
	backtrack = func(start int) {
		res = append(res, append([]int{}, paths...))
		used := make(map[int]bool, 0)

		for i := start; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}

			paths = append(paths, nums[i])

			backtrack(i + 1)

			paths = paths[:len(paths)-1]
			used[nums[i]] = true
		}
	}

	sort.Ints(nums)
	backtrack(0)

	return res
}
