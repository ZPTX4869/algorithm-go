package tree

import (
	"algorithm-go/structure/binarytree"
	"testing"
)

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"case1",
			args{
				root: binarytree.FromSlice([]int{3, 9, 20, null, null, 15, 7}).Root,
			},
			true,
		},
		{
			"case2",
			args{
				root: binarytree.FromSlice([]int{1, 2, 2, 3, 3, null, null, 4, 4}).Root,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
