package sorts

import "testing"

func Test_minNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"case1",
			args{
				nums: []int{10, 2},
			},
			"102",
		},
		{
			"case2",
			args{
				nums: []int{3, 30, 34, 5, 9},
			},
			"3033459",
		},
		{
			"case3",
			args{
				nums: []int{824, 8247},
			},
			"8247824",
		},
		{
			"case4",
			args{
				nums: []int{1, 1, 3, 2},
			},
			"1123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumber(tt.args.nums); got != tt.want {
				t.Errorf("minNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
