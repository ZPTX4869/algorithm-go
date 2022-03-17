package lists

import (
	"algorithm-go/structure/linkedlist"
	"reflect"
	"testing"
)

func Test_getKthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
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
				head: linkedlist.FromSlice([]int{1, 2, 3, 4, 5}).Head,
				k: 2,
			},
			linkedlist.FromSlice([]int{4, 5}).Head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKthFromEnd2(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getKthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
