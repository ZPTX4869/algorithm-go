package backtracks

import "testing"

func Test_movingCount(t *testing.T) {
	type args struct {
		m int
		n int
		k int
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
				m: 2,
				n: 3,
				k: 1,
			},
			3,
		},
		{
			"case2",
			args{
				m: 3,
				n: 1,
				k: 0,
			},
			1,
		},
		{
			"case3",
			args{
				m: 3,
				n: 2,
				k: 17,
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movingCount2(tt.args.m, tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("movingCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
