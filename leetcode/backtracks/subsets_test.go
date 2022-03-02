package backtracks

import (
	"fmt"
	"testing"
)

func Test_subsets(t *testing.T) {
	input := []int{1, 2, 3}
	output := subsets(input)

	fmt.Println(output)
}
