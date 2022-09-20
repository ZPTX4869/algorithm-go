package arr

import "algorithm-go/util/maths"

// 寻找两个正序数组中的中位数：https://leetcode.cn/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	findKth := func(k int) float64 {
		idx1, idx2 := 0, 0
		res := 0

		for i := 1; i <= k; i++ {
			if idx1 >= len(nums1) {
				res = nums2[idx2+k-i]
				break
			}
			if idx2 >= len(nums2) {
				res = nums1[idx1+k-i]
				break
			}

			if nums1[idx1] < nums2[idx2] {
				res = nums1[idx1]
				idx1 += 1
			} else {
				res = nums2[idx2]
				idx2 += 1
			}
		}

		return float64(res)
	}

	m, n := len(nums1), len(nums2)
	mid := (m + n) / 2

	if (m+n)%2 == 1 {
		return findKth(mid + 1)
	} else {
		smaller := findKth(mid)
		bigger := findKth(mid + 1)
		return (smaller + bigger) / 2
	}
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	var findKth func(s1, s2 int, k int) float64
	findKth = func(s1, s2 int, k int) float64 {
		l1, l2 := len(nums1)-s1, len(nums2)-s2

		if l1 == 0 {
			return float64(nums2[s2+k-1])
		}

		if l1 > l2 {
			nums1, nums2 = nums2, nums1
			return findKth(s2, s1, k)
		}

		if k == 1 {
			return float64(maths.Min(nums1[s1], nums2[s2]))
		}

		idx1 := s1 + maths.Min(l1, k/2) - 1
		idx2 := s2 + maths.Min(l2, k/2) - 1

		if nums1[idx1] < nums2[idx2] {
			return findKth(idx1+1, s2, k-maths.Min(l1, k/2))
		} else {
			return findKth(s1, idx2+1, k-maths.Min(l2, k/2))
		}
	}

	m, n := len(nums1), len(nums2)
	mid := (m + n) / 2

	if (m+n)%2 == 1 {
		return findKth(0, 0, mid+1)
	} else {
		smaller := findKth(0, 0, mid)
		bigger := findKth(0, 0, mid+1)
		return (smaller + bigger) / 2
	}
}
