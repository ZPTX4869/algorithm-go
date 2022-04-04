package stack

import (
	"fmt"
	"testing"
)

func Test_evalRPN(t *testing.T) {
	tokens := []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(tokens))
}
