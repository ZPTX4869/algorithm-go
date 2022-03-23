package trees

import (
	"algorithm-go/structure/binarytree"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"case1",
			args{
				binarytree.FromSlice([]int{3, 9, 20, null, null, 15, 7}).Root,
			},
			3,
		},
		{
			"case2",
			args{
				binarytree.FromSlice([]int{1, null, 2}).Root,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
