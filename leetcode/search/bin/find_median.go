package bin

import (
	"algorithm-go/util/maths"
)

// Bingary search
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len := len(nums1) + len(nums2)
	mid := len / 2
	if len%2 == 1 {
		return float64(findKthELem(nums1, nums2, mid+1))
	} else {
		return float64((findKthELem(nums1, nums2, mid) + findKthELem(nums1, nums2, mid+1)) / 2)
	}
}

func findKthELem(nums1 []int, nums2 []int, k int) float64 {
	idx1, idx2 := 0, 0

	for {
		if idx1 == len(nums1) {
			return float64(nums2[idx2+k-1])
		}
		if idx2 == len(nums2) {
			return float64(nums1[idx1+k-1])
		}
		if k == 1 {
			return float64(maths.Min(nums1[idx1], nums2[idx2]))
		}

		half := k / 2
		newIdx1 := maths.Min(idx1+half-1, len(nums1)-1)
		newIdx2 := maths.Min(idx2+half-1, len(nums2)-1)
		pivot1, pivot2 := nums1[newIdx1], nums2[newIdx2]

		if pivot1 < pivot2 {
			k -= (newIdx1 - idx1 + 1)
			idx1 = newIdx1 + 1
		} else {
			k -= (newIdx2 - idx2 + 1)
			idx2 = newIdx2 + 1
		}
	}
}
