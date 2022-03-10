package offer

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
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
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			2,
		},
		{
			"case2",
			args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 6,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search2(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
