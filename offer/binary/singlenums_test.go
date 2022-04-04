package binary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_singleNumbers(t *testing.T) {
	var nums []int
	var res []int

	nums = []int{4, 1, 4, 6}
	res = singleNumbers(nums)
	t.Run("case1", func(t *testing.T) {
		assert.Equal(t, []int{6, 1}, res)
	})

	nums = []int{1, 2, 5, 2}
	res = singleNumbers(nums)
	t.Run("case2", func(t *testing.T) {
		assert.Equal(t, []int{1, 5}, res)
	})
}
