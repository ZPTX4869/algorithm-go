package lists

import (
	"algorithm-go/structure/linkedlist"
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			"case1",
			args{
				head: linkedlist.FromSlice([]int{4, 5, 1, 9}).Head,
				val: 5,
			},
			linkedlist.FromSlice([]int{4, 1, 9}).Head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNode(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
