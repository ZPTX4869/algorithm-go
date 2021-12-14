package trees

import (
	"fmt"
	"learn-algorithm/structure/btree"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	vals := []int{5, 1, 4, null, null, 3, 6}
	tree := btree.FromSlice(vals)

	result := isValidBST(tree.Root)
	fmt.Println(result)
}
