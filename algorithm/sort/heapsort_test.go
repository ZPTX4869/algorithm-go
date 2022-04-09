package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestHeapSort_int(t *testing.T) {
	vals := []int{5, 4, 3, 2, 1}
	HeapSort(vals)

	assert.Equal(t, []int{1, 2, 3, 4, 5}, vals)
}

func TestHeapSort_float(t *testing.T) {
	vals := []int{3.0, 2.0, 1.0}
	HeapSort(vals)

	assert.Equal(t, []int{1.0, 2.0, 3.0}, vals)
}

func TestHeapSort_byte(t *testing.T) {
	vals := []int{'c', 'b', 'a'}
	HeapSort(vals)

	assert.Equal(t, []int{'a', 'b', 'c'}, vals)
}