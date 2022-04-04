package recur

import (
	"algorithm-go/structure/binarytree"
	"fmt"
	"testing"
)

func Test_generateTrees(t *testing.T) {
	roots := generateTrees(3)
	for _, r := range roots {
		result := binarytree.LevelTraverse(r)
		fmt.Println(result)
	}
	t.Fatal()
}
