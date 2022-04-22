package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeap(t *testing.T) {
	heap := NewMin[int]()

	heap.Push(5)
	heap.Push(3)
	heap.Push(1)

	assert.Equal(t, heap.Len(), uint(3))
	assert.Equal(t, heap.Empty(), false)

	assert.Equal(t, heap.Pop(), 1)
	assert.Equal(t, heap.Pop(), 3)
	assert.Equal(t, heap.Pop(), 5)
}

func TestMaxHeap(t *testing.T) {
	heap := NewMax[int]()

	heap.Push(5)
	heap.Push(3)
	heap.Push(1)

	assert.Equal(t, heap.Len(), uint(3))
	assert.Equal(t, heap.Empty(), false)

	assert.Equal(t, heap.Pop(), 5)
	assert.Equal(t, heap.Pop(), 3)
	assert.Equal(t, heap.Pop(), 1)
}
