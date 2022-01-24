package sorts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_quickSort(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1}
	QuickSort(nums)

	assert.Equal(t, []int{1, 2, 3, 4, 5}, nums)
}
