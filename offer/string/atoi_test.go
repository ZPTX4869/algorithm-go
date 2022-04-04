package string

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
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
				s: "9223372036854775808",
			},
			2147483647,
		},
		{
			"case2",
			args{
				s: "   -42",
			},
			-42,
		},
		{
			"case3",
			args{
				s: "4193 with words",
			},
			4193,
		},
		{
			"case4",
			args{
				s: "00000-42a1234",
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
