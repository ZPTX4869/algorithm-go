package list

import (
	"algorithm-go/structure/linkedlist"
	"fmt"
	"testing"
)

func Test_hasCycle(t *testing.T) {
	list := linkedlist.FromSlice([]int{3, 2, 0, -4})
	list.Tail.Next = list.Head.Next
	fmt.Println(hasCycle(list.Head))
}
