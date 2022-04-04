package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mergeSort(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1}
	MergeSort(nums)

	assert.Equal(t, []int{1, 2, 3, 4, 5}, nums)
}
