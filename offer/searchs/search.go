package searchs

func search(nums []int, target int) int {
	var res int

	for i, v := range nums {
		if v == target {
			res = 1
			for j := i + 1; j < len(nums) && nums[j] == target; j++ {
				res += 1
			}
			break
		}
	}

	return res
}

func search2(nums []int, target int) int {
	var res int
	var binarySearch func(s, e int)

	binarySearch = func(s, e int) {
		if s == e {
			res = 0
			return
		}

		leftIdx := (e-s)/2 + s
		rightIdx := leftIdx + 1

		if target > nums[leftIdx] {
			binarySearch(rightIdx, e)
		} else if target < nums[leftIdx] {
			binarySearch(s, leftIdx)
		} else {
			for rightIdx < len(nums) && nums[rightIdx] == target {
				rightIdx++
			}

			for leftIdx >= 0 && nums[leftIdx] == target {
				leftIdx--
			}

			res = rightIdx - leftIdx - 1
		}
	}

	binarySearch(0, len(nums))

	return res
}

func search3(nums []int, target int) int {
	var res int
	s, e := 0, len(nums)

	for s != e {
		mid := (e-s)/2 + s
		
		if target < nums[mid] {
			e = mid
		} else if target > nums[mid] {
			s = mid + 1
		} else {
			l := mid
			for l >= 0 && nums[l] == target {
				l--
			}

			r := mid+1
			for r < len(nums) && nums[r] == target {
				r++
			}

			res = r - l - 1
			break
		}
	}

	return res
}