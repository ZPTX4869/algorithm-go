package searchs

import "testing"

func Test_missingNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {
		// 	"case1",
		// 	args{
		// 		nums: []int{0, 1, 3},
		// 	},
		// 	2,
		// },
		{
			"case2",
			args{
				nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 9},
			},
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber2(tt.args.nums); got != tt.want {
				t.Errorf("missingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
