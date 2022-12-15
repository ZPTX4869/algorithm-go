package list

import (
	"algorithm-go/structure/linkedlist"
	"reflect"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case1",
			args: args{
				linkedlist.FromSlice([]int{1, 2, 3, 4}).Head,
			},
			want: linkedlist.FromSlice([]int{2, 1, 4, 3}).Head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swapPairs2(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case1",
			args: args{
				linkedlist.FromSlice([]int{1, 2, 3, 4}).Head,
			},
			want: linkedlist.FromSlice([]int{2, 1, 4, 3}).Head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs2() = %v, want %v", got, tt.want)
			}
		})
	}
}
