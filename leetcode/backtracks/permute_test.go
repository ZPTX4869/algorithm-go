package backtracks

import (
	"fmt"
	"testing"
)

func Test_permute(t *testing.T) {
	in := []int{1, 2, 3}
	out := permute(in)

	fmt.Println(out)
}
