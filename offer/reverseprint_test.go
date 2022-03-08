package offer

import (
	"algorithm-go/structure/linkedlist"
	"reflect"
	"testing"
)

func Test_reversePrint(t *testing.T) {
	type args struct {
		head *ListNode
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
				head: linkedlist.FromSlice([]int{1, 3, 2}).Head,
			},
			[]int{2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrint2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversePrint() = %v, want %v", got, tt.want)
			}
		})
	}
}
