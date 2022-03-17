package lists

import (
	"algorithm-go/structure/linkedlist"
	"fmt"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	h1 := linkedlist.FromSlice([]int{1, 3}).Head
	h2 := linkedlist.FromSlice([]int{2, 4}).Head

	// node := &ListNode {
	// 	Val: 5,
	// 	Next: nil,
	// }
	// h1.Next = node
	// h2.Next = node

	fmt.Println(getIntersectionNode2(h1, h2))
}
