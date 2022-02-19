package dynamics

import "testing"

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
