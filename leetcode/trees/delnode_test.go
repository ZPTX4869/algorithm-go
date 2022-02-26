package trees

import (
	"algorithm-go/structure/binarytree"
	"fmt"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	vals := []int{5,3,6,2,4,null,7}
	tree := binarytree.FromSlice(vals)
	fmt.Println(binarytree.LevelTraverse(tree.Root))

	tree.Root = deleteNode(tree.Root, 3)
	fmt.Println(binarytree.LevelTraverse(tree.Root))
}
