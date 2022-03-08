package offer

import (
	"fmt"
	"testing"
)

func TestMinStack(t *testing.T) {
	minStk := NewMinStack()
	minStk.Push(-2)
	minStk.Push(0)
	minStk.Push(-3)
	fmt.Println(minStk.GetMin())
	minStk.Pop()
	fmt.Println(minStk.Top())
	fmt.Println(minStk.GetMin())
}
