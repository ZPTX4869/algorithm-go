package list

import (
	"fmt"
	"testing"
)

func Test_copyRandomList(t *testing.T) {
	node0 := &Node{Val: 7}
	node1 := &Node{Val: 13}
	node2 := &Node{Val: 11}
	node3 := &Node{Val: 10}
	node4 := &Node{Val: 1}

	linkNode(node0, node1, node2, node3, node4)

	node0.Random = nil
	node1.Random = node0
	node2.Random = node4
	node3.Random = node2
	node4.Random = node0

	head := copyRandomList(node0)
	for head != nil {
		if head.Random != nil {
			fmt.Printf("[%v, %v]", head.Val, head.Random.Val)
		} else {
			fmt.Printf("[%v, nil]", head.Val)
		}

		head = head.Next
	}
}

func linkNode(nodes ...*Node) {
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
}
