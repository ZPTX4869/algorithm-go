package tree

import (
	"algorithm-go/structure/binarytree"
	"testing"
)

func Test_isSubStructure(t *testing.T) {
	type args struct {
		A *TreeNode
		B *TreeNode
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
				A: binarytree.FromSlice([]int{3, 4, 5, 1, 2}).Root,
				B: binarytree.FromSlice([]int{4, 1}).Root,
			},
			true,
		},
		{
			"case2",
			args{
				A: binarytree.FromSlice([]int{4, 2, 3, 4, 5, 6, 7, 8, 9}).Root,
				B: binarytree.FromSlice([]int{4, 8, 9}).Root,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubStructure2(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("isSubStructure() = %v, want %v", got, tt.want)
			}
		})
	}
}
