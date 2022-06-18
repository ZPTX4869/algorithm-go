package tree

import (
	"algorithm-go/structure/binarytree"
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{
			"case1",
			args{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			binarytree.FromSlice([]int{3, 9, 20, null, null, 15, 7}).Root,
		},
		{
			"case2",
			args{
				preorder: []int{1, 2},
				inorder:  []int{2, 1},
			},
			binarytree.FromSlice([]int{1, 2}).Root,
		},
		{
			"case3",
			args{
				preorder: []int{1, 2},
				inorder:  []int{1, 2},
			},
			binarytree.FromSlice([]int{1, null, 2}).Root,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
