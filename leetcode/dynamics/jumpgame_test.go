package dynamics

import (
	"testing"
)

func Test_canJump(t *testing.T) {
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
			name: "demo1",
			args: args{
				[]int{2, 3, 1, 1, 4},
			},
			want: true,
		},
		{
			name: "demo2",
			args: args{
				[]int{3, 2, 1, 0, 4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "demo1",
			args: args{
				[]int{2, 3, 1, 1, 4},
			},
			want: 2,
		},
		{
			name: "demo2",
			args: args{
				[]int{2, 3, 0, 1, 4},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
