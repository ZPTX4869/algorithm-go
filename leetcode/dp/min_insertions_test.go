package dp

import "testing"

func Test_minInsertions(t *testing.T) {
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
				s: "zzazz",
			},
			want: 0,
		},
		{
			name: "case2",
			args: args{
				s: "mbadm",
			},
			want: 2,
		},
		{
			name: "case3",
			args: args{
				s: "leetcode",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minInsertions(tt.args.s); got != tt.want {
				t.Errorf("minInsertions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minInsertions2(t *testing.T) {
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
				s: "zzazz",
			},
			want: 0,
		},
		{
			name: "case2",
			args: args{
				s: "mbadm",
			},
			want: 2,
		},
		{
			name: "case3",
			args: args{
				s: "leetcode",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minInsertions2(tt.args.s); got != tt.want {
				t.Errorf("minInsertions2() = %v, want %v", got, tt.want)
			}
		})
	}
}
