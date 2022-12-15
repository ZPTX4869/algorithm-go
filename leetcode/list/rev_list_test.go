package list

import (
	"algorithm-go/structure/linkedlist"
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
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
				head: linkedlist.FromSlice([]int{1, 2, 3, 4, 5}).Head,
            },
            want: linkedlist.FromSlice([]int{5, 4, 3, 2, 1}).Head,
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseList2(t *testing.T) {
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
				head: linkedlist.FromSlice([]int{1, 2, 3, 4, 5}).Head,
            },
            want: linkedlist.FromSlice([]int{5, 4, 3, 2, 1}).Head,
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList2() = %v, want %v", got, tt.want)
			}
		})
	}
}
