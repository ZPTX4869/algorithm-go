package search

import "testing"

func TestLeftBound(t *testing.T) {
	type args struct {
		vals []int
		tar  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"case1",
			args{
				vals: []int{1, 2, 2, 4, 5},
				tar:  2,
			},
			1,
		},
		{
			"case2",
			args{
				vals: []int{1, 1, 3, 5, 7},
				tar:  0,
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LeftBound(tt.args.vals, tt.args.tar); got != tt.want {
				t.Errorf("LeftBound() = %v, want %v", got, tt.want)
			}
		})
	}
}
