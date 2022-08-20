package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyCircularQueue(t *testing.T) {
	circularQueue := Constructor(3)
	
	circularQueue.EnQueue(1)
	circularQueue.EnQueue(2)
	circularQueue.EnQueue(3)
	circularQueue.EnQueue(4)

	assert.Equal(t, circularQueue.Rear(), 3)
	assert.True(t, circularQueue.IsFull())
	assert.True(t, circularQueue.DeQueue())
	assert.True(t, circularQueue.EnQueue(4))
	assert.Equal(t, circularQueue.Rear(), 4)
}
