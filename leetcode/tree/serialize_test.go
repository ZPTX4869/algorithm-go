package tree

import (
	"algorithm-go/structure/binarytree"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodec(t *testing.T) {

	root := binarytree.FromSlice([]int{1, 2, 3, null, null, 4, 5}).Root

	ser := Constructor()
	deser := Constructor()
	data := ser.serialize(root)
	nRoot := deser.deserialize(data)

	assert.True(t, reflect.DeepEqual(root, nRoot))
}

func TestCodec2(t *testing.T) {
	root := binarytree.FromSlice([]int{1, 2, 3, null, null, 4, 5}).Root

	ser := Constructor()
	deser := Constructor()
	data := ser.serialize2(root)
	nRoot := deser.deserialize2(data)

	assert.Equal(t, root, nRoot)
}
