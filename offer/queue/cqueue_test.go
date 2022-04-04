package queue

import (
	"fmt"
	"testing"
)

func TestCQueue(t *testing.T) {
	q := NewCQueue()
	q.DeleteHead()

	q.AppendTail(3)
	q.AppendTail(4)
	q.AppendTail(5)

	q.DeleteHead()
	q.DeleteHead()

	fmt.Println(q)
}
