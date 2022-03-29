package binarys

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func Test_singleNumbers2(t *testing.T) {
	var nums []int
	var res int

	nums = []int{3, 4, 3, 3}
	res = singleNumber(nums)
	t.Run("case1", func(t *testing.T) {
		assert.Equal(t, 4, res)
	})

	nums = []int{9, 1, 7, 9, 7, 9, 7}
	res = singleNumber(nums)
	t.Run("case1", func(t *testing.T) {
		assert.Equal(t, 1, res)
	})

}
