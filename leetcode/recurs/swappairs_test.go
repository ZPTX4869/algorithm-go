package recurs

import (
	"algorithm-go/structure/linkedlist"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_swapPairs(t *testing.T) {
	list := linkedlist.FromSlice([]int{1, 2, 3, 4})
	list.Head = swapPairs(list.Head)
	res := list.Traverse()

	assert.Equal(t, []int{2, 1, 4, 3}, res)
}
