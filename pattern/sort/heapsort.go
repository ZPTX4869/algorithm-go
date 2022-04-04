package sort

func HeapSort(nums []int) {
	buildHeap(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, 0, i)
	}
	return
}

func buildHeap(nums []int) {
	size := len(nums)
	for i := size/2 - 1; i >= 0; i-- {
		heapify(nums, i, size)
	}
}

func heapify(nums []int, i, size int) {
	left, right := i*2+1, i*2+2
	max := i

	if left < size && nums[max] < nums[left] {
		max = left
	}

	if right < size && nums[max] < nums[right] {
		max = right
	}

	if max != i {
		nums[i], nums[max] = nums[max], nums[i]
		heapify(nums, max, size)
	}
}
