package sort

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left, right int) {
	if left < right {
		pivotIdx := partition(nums, left, right)
		quickSort(nums, left, pivotIdx-1)
		quickSort(nums, pivotIdx+1, right)
	}
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] < pivot {
			swap(nums, i, j)
			i++
		}
	}
	swap(nums, i, right)
	return i
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
