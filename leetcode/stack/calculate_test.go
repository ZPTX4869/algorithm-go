package stack

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {
		// 	name: "case1",
		// 	args: args{
		// 		s: "1 + 1",
		// 	},
		// 	want: 2,
		// },
		// {
		// 	name: "case2",
		// 	args: args{
		// 		s: " 2-1 + 2 ",
		// 	},
		// 	want: 3,
		// },
		// {
		// 	name: "case3",
		// 	args: args{
		// 		s: "(1+2)*3",
		// 	},
		// 	want: 9,
		// },
		{
			name: "case",
			args: args{
				s: "2147483647",
			},
			want: 2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
