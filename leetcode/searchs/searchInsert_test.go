package searchs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_searchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	res := searchInsert(nums, 7)

	assert.Equal(t, 4, res)
}
