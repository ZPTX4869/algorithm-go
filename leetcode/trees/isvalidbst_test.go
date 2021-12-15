package trees

import (
	"algorithm-go/structure/binarytree"
	"fmt"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	vals := []int{5, 1, 4, null, null, 3, 6}
	tree := binarytree.FromSlice(vals)

	result := isValidBST(tree.Root)
	fmt.Println(result)
}
