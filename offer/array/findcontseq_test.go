package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findContinuousSequence(t *testing.T) {
	type args struct {
		target int
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
				9,
			},
			[][]int{
				{2, 3, 4},
				{4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findContinuousSequence(tt.args.target), "findContinuousSequence(%v)", tt.args.target)
		})
	}
}
