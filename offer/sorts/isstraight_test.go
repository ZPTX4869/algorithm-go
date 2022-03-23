package sorts

import "testing"

func Test_isStraight(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"case1",
			args{
				nums: []int{1, 2, 3, 4, 5},
			},
			true,
		},
		{
			"case2",
			args{
				nums: []int{0, 0, 1, 2, 5},
			},
			true,
		},
		{
			"case3",
			args{
				nums: []int{1, 2, 12, 0, 3},
			},
			false,
		},
		{
			"case4",
			args{
				nums: []int{0, 0, 2, 2, 5},
			},
			false,
		},
		{
			"case5",
			args{
				nums: []int{10, 11, 0, 12, 6},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStraight(tt.args.nums); got != tt.want {
				t.Errorf("isStraight() = %v, want %v", got, tt.want)
			}
		})
	}
}
