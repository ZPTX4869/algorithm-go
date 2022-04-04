package backtrack

import (
	"fmt"
	"testing"
)

func Test_subsets(t *testing.T) {
	input := []int{1, 2, 3}
	output := subsets(input)

	fmt.Println(output)
}

func Test_subsetsWithDup(t *testing.T) {
	input := []int{1, 2, 2}
	output := subsetsWithDup(input)

	fmt.Println(output)
}
