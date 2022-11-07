package list

import (
	"algorithm-go/structure/linkedlist"
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// : Add test cases.
		{
			name: "case1",
			args: args{
				head: linkedlist.FromSlice([]int{1, 2, 3, 3, 4, 4, 5}).Head,
			},
			want: linkedlist.FromSlice([]int{1, 2, 5}).Head,
		},
		{
			name: "case2",
			args: args{
				head: linkedlist.FromSlice([]int{1, 1, 1, 2, 3}).Head,
			},
			want: linkedlist.FromSlice([]int{2, 3}).Head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteDuplicates2(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// : Add test cases.
		{
			name: "case1",
			args: args{
				head: linkedlist.FromSlice([]int{1, 2, 3, 3, 4, 4, 5}).Head,
			},
			want: linkedlist.FromSlice([]int{1, 2, 5}).Head,
		},
		{
			name: "case2",
			args: args{
				head: linkedlist.FromSlice([]int{1, 1, 1, 2, 3}).Head,
			},
			want: linkedlist.FromSlice([]int{2, 3}).Head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates2() = %v, want %v", got, tt.want)
			}
		})
	}
}
