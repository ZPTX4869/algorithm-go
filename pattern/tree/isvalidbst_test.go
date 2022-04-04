package tree

import (
	"algorithm-go/structure/binarytree"
	"testing"
)

func Test_isValidBST(t *testing.T) {
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
				binarytree.FromSlice([]int{2, 1, 3}).Root,
			},
			true,
		},
		{
			"case2",
			args{
				binarytree.FromSlice([]int{5, 1, 4, null, null, 3, 6}).Root,
			},
			false,
		},
		{
			"case3",
			args{
				binarytree.FromSlice([]int{2, 2, 2}).Root,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidBST2(t *testing.T) {
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
				binarytree.FromSlice([]int{2, 1, 3}).Root,
			},
			true,
		},
		{
			"case2",
			args{
				binarytree.FromSlice([]int{5, 1, 4, null, null, 3, 6}).Root,
			},
			false,
		},
		{
			"case3",
			args{
				binarytree.FromSlice([]int{2, 2, 2}).Root,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST2(tt.args.root); got != tt.want {
				t.Errorf("isValidBST2() = %v, want %v", got, tt.want)
			}
		})
	}
}
