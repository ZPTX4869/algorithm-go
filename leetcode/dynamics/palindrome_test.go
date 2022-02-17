package dynamics

import (
	"reflect"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "demo1",
			args: args{
				"babad",
			},
			want: "bab",
		},
		{
			name: "demo2",
			args: args{
				"cbbd",
			},
			want: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
		{
			name: "demo1",
			args: args{
				"aab",
			},
			want: [][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
		{
			name: "demo2",
			args: args{
				"a",
			},
			want: [][]string{
				{"a"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minCut(t *testing.T) {
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
			name: "demo1",
			args: args{
				"aab",
			},
			want: 1,
		},
		{
			name: "demo2",
			args: args{
				"a",
			},
			want: 0,
		},
		{
			name: "demo3",
			args: args{
				"ab",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCut(tt.args.s); got != tt.want {
				t.Errorf("minCut() = %v, want %v", got, tt.want)
			}
		})
	}
}
