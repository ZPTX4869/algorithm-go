package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort_int(t *testing.T) {
	vals := []int{5, 4, 3, 2, 1}
	BubbleSort(vals)

	assert.Equal(t, []int{1, 2, 3, 4, 5}, vals)
}

func TestBubbleSort_float(t *testing.T) {
	vals := []int{3.0, 2.0, 1.0}
	BubbleSort(vals)

	assert.Equal(t, []int{1.0, 2.0, 3.0}, vals)
}

func TestBubbleSort_byte(t *testing.T) {
	vals := []int{'c', 'b', 'a'}
	BubbleSort(vals)

	assert.Equal(t, []int{'a', 'b', 'c'}, vals)
}
