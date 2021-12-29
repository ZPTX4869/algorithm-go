package stacks

import (
	"fmt"
	"testing"
)

func Test_decodeString(t *testing.T) {
	s := "10[leetcode]"
	fmt.Println(decodeString(s))
}
