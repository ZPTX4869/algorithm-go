package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseWords(t *testing.T) {
	s := "the bule is sky"
	res := reverseWords2(s)

	assert.Equal(t, "sky is bule the", res)
}



