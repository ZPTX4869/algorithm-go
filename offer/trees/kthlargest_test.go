package trees

import (
	"algorithm-go/structure/binarytree"
	"testing"
)

func Test_kthLargest(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
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
				root: binarytree.FromSlice([]int{3,1,4,null,2}).Root,
				k: 1,
			},
			4,
		},
		{
			"case2",
			args{
				root: binarytree.FromSlice([]int{5, 3, 6, 2, 4, null, null, 1}).Root,
				k:    3,
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthLargest2(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
