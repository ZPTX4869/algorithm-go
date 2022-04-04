package list

import (
	"algorithm-go/structure/linkedlist"
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	list := linkedlist.FromSlice([]int{1, 1, 2, 1})
	fmt.Println(isPalindrome(list.Head))
}
