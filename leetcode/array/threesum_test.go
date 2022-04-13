package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_threeSum(t *testing.T) {
	var nums []int
	var res [][]int

	t.Run("case1", func(t *testing.T) {
		nums = []int{-1, 0, 1, 2, -1, -4}

		res = threeSum(nums)
		assert.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, res)
	})

	t.Run("case2", func(t *testing.T) {
		nums = []int{}

		res = threeSum(nums)
		assert.Equal(t, [][]int{}, res)
	})

	t.Run("case3", func(t *testing.T) {
		nums = []int{0, 0, 0}

		res = threeSum(nums)
		assert.Equal(t, [][]int{{0, 0, 0}}, res)
	})

	t.Run("case4", func(t *testing.T) {
		nums = []int{0, 0, 0, 0}

		res = threeSum(nums)
		assert.Equal(t, [][]int{{0, 0, 0}}, res)
	})
}
