package searchs

func missingNumber(nums []int) int {
	var res int
	var binarySearch func(s, e int)

	binarySearch = func(s, e int) {
		if s == e {
			res = s
			return
		}

		mid := (e-s)/2 + s

		if nums[mid] == mid {
			binarySearch(mid+1, e)
		} else {
			binarySearch(s, mid)
		}
	}

	binarySearch(0, len(nums))

	return res
}

func missingNumber2(nums []int) int {
	s, e := 0, len(nums)
	for s != e {
		mid := (e-s)/2 + s
		if nums[mid] == mid {
			s = mid+1
		} else {
			e = mid
		}
	}

	return s
}