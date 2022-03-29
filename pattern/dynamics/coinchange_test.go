package dynamics

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
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
				coins:  []int{1, 2, 5},
				amount: 11,
			},
			3,
		},
		{
			"case2",
			args{
				coins:  []int{2},
				amount: 3,
			},
			-1,
		},
		{
			"case3",
			args{
				coins:  []int{1},
				amount: 0,
			},
			0,
		},
		{
			"case4",
			args{
				coins:  []int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422},
				amount: 9864,
			},
			24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
