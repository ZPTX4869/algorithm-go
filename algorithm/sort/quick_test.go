package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort_int(t *testing.T) {
	vals := []int{5, 4, 3, 2, 1}
	QuickSort(vals)

	assert.Equal(t, []int{1, 2, 3, 4, 5}, vals)
}

func TestQuickSort_float(t *testing.T) {
	vals := []int{3.0, 2.0, 1.0}
	QuickSort(vals)

	assert.Equal(t, []int{1.0, 2.0, 3.0}, vals)
}

func TestQuickSort_byte(t *testing.T) {
	vals := []int{'c', 'b', 'a'}
	QuickSort(vals)

	assert.Equal(t, []int{'a', 'b', 'c'}, vals)
}
