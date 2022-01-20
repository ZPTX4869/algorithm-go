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

func Test_findMin2(t *testing.T) {
	nums := []int{10, 1, 10, 10, 10}
	res := findMin2(nums)

	assert.Equal(t, 1, res)
}
