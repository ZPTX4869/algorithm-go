package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_search(t *testing.T) {
	nums := []int{3, 5, 1}
	res := search(nums, 1)

	assert.Equal(t, 2, res)
}

func Test_search2(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 13, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	res := search2(nums, 13)

	assert.Equal(t, true, res)
}
