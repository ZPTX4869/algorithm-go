package searchs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMin(t *testing.T) {
	nums := []int{2, 3, 4, 5, 1}
	res := findMin(nums)

	assert.Equal(t, 1, res)
}
