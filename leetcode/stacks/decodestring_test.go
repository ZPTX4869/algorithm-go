package stacks

import (
	"fmt"
	"testing"
)

func Test_decodeString(t *testing.T) {
	s := "3[a]2[bc]"
	fmt.Println(decodeString(s))
}
