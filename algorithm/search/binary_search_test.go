package search

import "testing"

func TestBinarySearch(t *testing.T) {
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
				vals: []int{1, 2, 3, 4, 5, 6, 7},
				tar:  3,
			},
			2,
		},
		{
			"case2",
			args{
				vals: []int{1, 3, 5, 7},
				tar:  3,
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.vals, tt.args.tar); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
