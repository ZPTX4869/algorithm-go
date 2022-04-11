package util

import (
	"os"
	"testing"
)

func TestCheckOrPanic(t *testing.T) {
	_, err := os.ReadFile("../nonexist.txt")
	CheckOrPanic(err)
}

func TestCheckOrFatal(t *testing.T) {
	_, err := os.ReadFile("../nonexist.txt")
	CheckOrFatal(err)
}
