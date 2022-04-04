package backtrack

import (
	"fmt"
	"testing"
)

func Test_permute(t *testing.T) {
	in := []int{1, 2, 3}
	out := permute(in)

	fmt.Println(out)
}

func Test_permuteUnique(t *testing.T) {
	in := []int{1, 1, 2}
	out := permuteUnique(in)

	fmt.Println(out)
}
