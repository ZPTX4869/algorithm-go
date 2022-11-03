package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nextPermutation(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		nums := []int{1, 2, 3}
		expected := []int{1, 3, 2}

		nextPermutation(nums)
		assert.Equal(t, expected, nums)
	})

	t.Run("case2", func(t *testing.T) {
		nums := []int{3, 2, 1}
		expected := []int{3, 2, 1}

		nextPermutation(nums)
		assert.Equal(t, expected, nums)
	})
}
