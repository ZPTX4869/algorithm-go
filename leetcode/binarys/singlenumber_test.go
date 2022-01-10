package binarys

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	nums := []int{1, 2, 1, 2, -2}

	res := singleNumber(nums)
	assert.Equal(t, -2, res)
}

func Test_singleNumber2(t *testing.T) {
	nums := []int{2, 2, 3, 2}

	res := singleNumber2(nums)
	assert.Equal(t, 3, res)
}

func Test_singleNumber3(t *testing.T) {
	nums := []int{1, 1, 2, 3, 3, 4}

	res := singleNumber3(nums)
	assert.Equal(t, []int{4, 2}, res)
}
