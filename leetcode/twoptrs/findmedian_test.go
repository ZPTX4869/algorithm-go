package twoptrs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMedianSortedArrays(t *testing.T) {
	var nums1, nums2 []int

	t.Run("case1", func(t *testing.T) {
		nums1 = []int{1, 3}
		nums2 = []int{2}

		res := findMedianSortedArrays(nums1, nums2)
		assert.Equal(t, float64(2), res)
	})

	t.Run("case2", func(t *testing.T) {
		nums1 = []int{1, 3}
		nums2 = []int{2, 4}

		res := findMedianSortedArrays(nums1, nums2)
		assert.Equal(t, float64(2.5), res)
	})

	t.Run("case3", func(t *testing.T) {
		nums1 = []int{}
		nums2 = []int{2, 4}

		res := findMedianSortedArrays(nums1, nums2)
		assert.Equal(t, float64(3), res)
	})
}
