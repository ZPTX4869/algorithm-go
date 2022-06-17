package tree

import (
	"algorithm-go/structure/binarytree"
	"reflect"
	"testing"
)

func Test_invertTree(t *testing.T) {
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
				root: binarytree.FromSlice([]int{4, 2, 7, 1, 3, 6, 9}).Root,
			},
			binarytree.FromSlice([]int{4, 7, 2, 9, 6, 3, 1}).Root,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_invertTree2(t *testing.T) {
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
				root: binarytree.FromSlice([]int{4, 2, 7, 1, 3, 6, 9}).Root,
			},
			binarytree.FromSlice([]int{4, 7, 2, 9, 6, 3, 1}).Root,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
