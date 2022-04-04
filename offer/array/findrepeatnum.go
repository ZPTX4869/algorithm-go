package array

func findRepeatNumber(nums []int) int {
	visited := make(map[int]bool, 0)

	for _, num := range nums {
		if visited[num] {
			return num
		} 

		visited[num] = true
	}

	return nums[0]
}