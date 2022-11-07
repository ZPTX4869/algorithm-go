package graph

import "testing"

func Test_equationsPossible(t *testing.T) {
	type args struct {
		equations []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// : Add test cases.
		{
			name: "case1",
			args: args{
				equations: []string{"a!=i", "g==k", "k==j", "k!=i", "c!=e", "a!=e", "k!=a", "a!=g", "g!=c"}},
			want: true,
		},
		{
			name: "case1",
			args: args{
				equations: []string{"a==b", "b!=c", "c==a"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equationsPossible(tt.args.equations); got != tt.want {
				t.Errorf("equationsPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
