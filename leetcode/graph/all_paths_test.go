package graph

import (
	"reflect"
	"testing"
)

func Test_allPathsSourceTarget(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// : Add test cases.
		{
			"case1",
			args{
				graph: [][]int{
					{1, 2},
					{3},
					{3},
					{},
				},
			},
			[][]int{{0, 1, 3}, {0, 2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allPathsSourceTarget(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allPathsSourceTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
