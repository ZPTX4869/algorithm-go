package tree

import (
	"algorithm-go/structure/binarytree"
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"case1",
			args{
				root: binarytree.FromSlice([]int{3, 9, 20, null, null, 15, 7}).Root,
			},
			[]int{3, 9, 20, 15, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_levelOrder2(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			"case1",
			args{
				root: binarytree.FromSlice([]int{3, 9, 20, null, null, 15, 7}).Root,
			},
			[][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_levelOrder3(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			"case1",
			args{
				root: binarytree.FromSlice([]int{3, 9, 20, null, null, 15, 7}).Root,
			},
			[][]int{
				{3},
				{20, 9},
				{15, 7},
			},
		},
		{
			"case2",
			args{
				root: binarytree.FromSlice([]int{1, 2, 3, 4, null, null, 5}).Root,
			},
			[][]int{
				{1},
				{3, 2},
				{4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder3(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder3() = %v, want %v", got, tt.want)
			}
		})
	}
}
