package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	var nums []int
	var res int

	t.Run("case1", func(t *testing.T) {
		nums = []int{1, 2, 3}
		
		res = Sum(nums)
		assert.Equal(t, 6, res)
	})
}
