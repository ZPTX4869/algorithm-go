package dfs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_letterCombinations(t *testing.T) {
	var digits string
	var res []string

	t.Run("case1", func(t *testing.T) {
		digits = "23"
		res = letterCombinations(digits)

		assert.Equal(t, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, res)
	})

	t.Run("case2", func(t *testing.T) {
		digits = ""
		res = letterCombinations(digits)

		assert.Equal(t, []string{}, res)
	})
}
