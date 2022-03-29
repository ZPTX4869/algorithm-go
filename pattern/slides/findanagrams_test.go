package slides

import (
	"reflect"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"case1",
			args{
				s: "cbaebabacd",
				p: "abc",
			},
			[]int{0, 6},
		},
		{
			"case2",
			args{
				s: "abab",
				p: "ab",
			},
			[]int{0, 1, 2},
		},
		{
			"case3",
			args{
				s: "aa",
				p: "bb",
			},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagrams(tt.args.s, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
