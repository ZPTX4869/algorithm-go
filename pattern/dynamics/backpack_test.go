package dynamics

import (
	"testing"
)

func Test_backPack(t *testing.T) {
	type args struct {
		m int
		A []int
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
				m: 10,
				A: []int{3, 4, 8, 5},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backPack(tt.args.m, tt.args.A); got != tt.want {
				t.Errorf("backPack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backPackII(t *testing.T) {
	type args struct {
		m int
		A []int
		V []int
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
				m: 10,
				A: []int{2, 3, 5, 7},
				V: []int{1, 5, 2, 4},
			},
			want: 9,
		},
		{
			name: "case2",
			args: args{
				m: 300,
				A: []int{95,75,23,73,50,22,6,57,89,98},
				V: []int{89,59,19,43,100,72,44,16,7,64},
			},
			want: 388,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backPackII(tt.args.m, tt.args.A, tt.args.V); got != tt.want {
				t.Errorf("backPackII() = %v, want %v", got, tt.want)
			}
		})
	}
}
