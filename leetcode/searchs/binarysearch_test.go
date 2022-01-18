package searchs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_binarySearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9}
	res := binarySearch(nums, 2)

	assert.Equal(t, -1, res)
}

func Test_binarySearch2(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	res := binarySearch2(nums, 2)

	assert.Equal(t, -1, res)
}

func Test_binarySearch3(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	res := binarySearch3(nums, 2)

	assert.Equal(t, -1, res)
}
