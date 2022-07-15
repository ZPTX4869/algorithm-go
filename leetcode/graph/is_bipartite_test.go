package graph

import (
	"testing"
)

func Test_isBipartite(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				graph: [][]int{
					{1, 2, 3},
					{0, 2},
					{0, 1, 3},
					{0, 2},
				},
			},
			want: false,
		},
		{
			name: "case2",
			args: args{
				graph: [][]int{
					{1, 3},
					{0, 2},
					{1, 3},
					{0, 2},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBipartite(tt.args.graph); got != tt.want {
				t.Errorf("isBipartite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isBipartite_(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				graph: [][]int{
					{1, 2, 3},
					{0, 2},
					{0, 1, 3},
					{0, 2},
				},
			},
			want: false,
		},
		{
			name: "case2",
			args: args{
				graph: [][]int{
					{1, 3},
					{0, 2},
					{1, 3},
					{0, 2},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBipartite_(tt.args.graph); got != tt.want {
				t.Errorf("isBipartite_() = %v, want %v", got, tt.want)
			}
		})
	}
}
