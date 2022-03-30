package lru

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	n1 := NewNode(1, 1)
	n2 := NewNode(2, 2)
	n3 := NewNode(3, 3)

	l := NewList()

	l.PushFront(n1)
	l.PushFront(n2)
	l.PushFront(n3)
	fmt.Printf("h: (%d, %d), t: (%d, %d)\n", l.head.Key, l.head.Val, l.tail.Key, l.tail.Val)

	l.Remove(n2)
	fmt.Printf("h: (%d, %d), t: (%d, %d)\n", l.head.Key, l.head.Val, l.tail.Key, l.tail.Val)

	l.Remove(n1)
	fmt.Printf("h: (%d, %d), t: (%d, %d)\n", l.head.Key, l.head.Val, l.tail.Key, l.tail.Val)

	l.PopBack()
}