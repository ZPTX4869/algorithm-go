package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_majorityElement(t *testing.T) {
	var nums []int
	var res int

	nums = []int{2, 2, 1, 1, 1, 2, 2}
	res = majorityElement(nums)
	t.Run("case1", func(t *testing.T) {
		assert.Equal(t, 2, res)
	})
}

func Test_majorityElement2(t *testing.T) {
	var nums []int
	var res int

	nums = []int{2, 2, 1, 1, 1, 2, 2}
	res = majorityElement2(nums)
	t.Run("case1", func(t *testing.T) {
		assert.Equal(t, 2, res)
	})
}