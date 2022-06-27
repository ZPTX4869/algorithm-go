package tree

import (
	"algorithm-go/structure/binarytree"
	"reflect"
	"testing"
)

func Test_bstToGst(t *testing.T) {
	type args struct {
		root *TreeNode
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
				root: binarytree.FromSlice([]int{4, 1, 6, 0, 2, 5, 7, null, null, null, 3, null, null, null, 8}).Root,
			},
			binarytree.FromSlice([]int{30, 36, 21, 36, 35, 26, 15, null, null, null, 33, null, null, null, 8}).Root,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bstToGst(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bstToGst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertBST(t *testing.T) {
	type args struct {
		root *TreeNode
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
				root: binarytree.FromSlice([]int{4, 1, 6, 0, 2, 5, 7, null, null, null, 3, null, null, null, 8}).Root,
			},
			binarytree.FromSlice([]int{30, 36, 21, 36, 35, 26, 15, null, null, null, 33, null, null, null, 8}).Root,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertBST(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
