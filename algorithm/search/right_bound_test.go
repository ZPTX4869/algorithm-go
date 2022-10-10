package search

import "testing"

func TestRightBound(t *testing.T) {
	type args struct {
		vals []int
		tar  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{
				vals: []int{1, 2, 4, 4, 5},
				tar:  4,
			},
			3,
		},
		{
			"case2",
			args{
				vals: []int{1, 3, 5, 7, 7},
				tar:  4,
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RightBound(tt.args.vals, tt.args.tar); got != tt.want {
				t.Errorf("RightBound() = %v, want %v", got, tt.want)
			}
		})
	}
}
