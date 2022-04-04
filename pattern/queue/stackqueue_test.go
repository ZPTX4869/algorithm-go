package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStackQueue(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	assert.Equal(t, 1, queue.Peek())
	assert.Equal(t, 1, queue.Pop())
	assert.Equal(t, false, queue.Empty())
}
