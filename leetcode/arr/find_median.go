package arr

import "algorithm-go/util/maths"

// 寻找两个正序数组中的中位数：https://leetcode.cn/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	getKth := func(k int) float64 {
		idx1, idx2 := 0, 0
		kth := 0

		for i := 1; i <= k; i++ {
			if idx1 >= len(nums1) {
				kth = nums2[idx2+k-i]
				break
			}
			if idx2 >= len(nums2) {
				kth = nums1[idx1+k-i]
				break
			}

			if nums1[idx1] < nums2[idx2] {
				kth = nums1[idx1]
				idx1 += 1
			} else {
				kth = nums2[idx2]
				idx2 += 1
			}
		}

		return float64(kth)
	}

	m, n := len(nums1), len(nums2)
	left := getKth((m + n + 1) / 2)
	right := getKth((m + n + 2) / 2)

	return (left + right) / 2
}

func findMedianSortedArrays_(nums1 []int, nums2 []int) float64 {
	var getKth func(nums1 []int, s1 int, nums2 []int, s2 int, k int) float64
	getKth = func(nums1 []int, s1 int, nums2 []int, s2 int, k int) float64 {
		n1, n2 := len(nums1)-s1, len(nums2)-s2

		if n1 > n2 {
			return getKth(nums2, s2, nums1, s1, k)
		}

		if n1 == 0 {
			return float64(nums2[s2+k-1])
		}

		if k == 1 {
			return float64(maths.Min(nums1[s1], nums2[s2]))
		}

		i := s1 + maths.Min(n1, k/2) - 1
		j := s2 + maths.Min(n2, k/2) - 1

		if nums1[i] < nums2[j] {
			return getKth(nums1, i+1, nums2, s2, k-(i-s1+1))
		} else {
			return getKth(nums1, s1, nums2, j+1, k-(j-s2+1))
		}

	}

	m, n := len(nums1), len(nums2)
	left := getKth(nums1, 0, nums2, 0, (m+n+1)/2)
	right := getKth(nums1, 0, nums2, 0, (m+n+2)/2)

	return (left + right) / 2
}
