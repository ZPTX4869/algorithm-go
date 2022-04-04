package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_cloneGraph(t *testing.T) {
	node1 := &Node{
		Val:       1,
		Neighbors: make([]*Node, 0),
	}

	node2 := &Node{
		Val:       2,
		Neighbors: make([]*Node, 0),
	}

	node3 := &Node{
		Val:       3,
		Neighbors: make([]*Node, 0),
	}

	node1.Neighbors = append(node1.Neighbors, node2, node3)
	node2.Neighbors = append(node2.Neighbors, node1, node3)
	node3.Neighbors = append(node3.Neighbors, node1, node2)

	headCopy := cloneGraph(node1)
	assert.Equal(t, node1, headCopy)
}
