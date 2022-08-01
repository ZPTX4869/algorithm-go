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
		{
            name: "case1",
            args: args{
				s: "405+5",
            },
            want: 410,
        },
		{
            name: "case2",
            args: args{
				s: "(1+2)*3",
            },
            want: 9,
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

func Test_calculate_(t *testing.T) {
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
            name: "case1",
            args: args{
				s: "405+5",
            },
            want: 410,
        },
		{
            name: "case2",
            args: args{
				s: "(1+2)*3",
            },
            want: 9,
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate_(tt.args.s); got != tt.want {
				t.Errorf("calculate_() = %v, want %v", got, tt.want)
			}
		})
	}
}