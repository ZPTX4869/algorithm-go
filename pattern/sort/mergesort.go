package sort

func MergeSort(nums []int) {
	if len(nums) == 0 {
		return
	}

	mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, left, right int) {
	if left == right {
		return
	}

	mid := left + (right-left)/2
	mergeSort(nums, left, mid)
	mergeSort(nums, mid+1, right)
	merge(nums, left, mid, right)
}

func merge(nums []int, left, mid, right int) {
	res := make([]int, 0)

	i, j := left, mid+1
	for i <= mid && j <= right {
		if nums[left] <= nums[right] {
			res = append(res, nums[i])
			i += 1
		} else {
			res = append(res, nums[j])
			j += 1
		}
	}

	if i > mid {
		res = append(res, nums[j:right+1]...)
	} else {
		res = append(res, nums[i:mid+1]...)
	}

	copy(nums[left:right+1], res)
}
