package str

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
		// : Add test cases.
		{
			name: "case1",
			args: args{
				s: "   -42",
			},
			want: -42,
		},
		{
			name: "case2",
			args: args{
				s: "4193 with words",
			},
			want: 4193,
		},
		{
			name: "case3",
			args: args{
				s: "words and 987",
			},
			want: 0,
		},
		{
			name: "case4",
			args: args{
				s: "9223372036854775808",
			},
			want: 2147483647,
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
