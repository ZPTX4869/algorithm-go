package dp

import (
	"algorithm-go/structure/binarytree"
	"testing"
)

func Test_rob(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// : Add test cases.
		{
			name: "case1",
			args: args{
				root: binarytree.FromSlice([]int{3, 2, 3, null, 3, null, 1}).Root,
			},
			want: 7,
		},
		{
			name: "case2",
			args: args{
				root: binarytree.FromSlice([]int{3, 4, 5, 1, 3, null, 1}).Root,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.root); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
