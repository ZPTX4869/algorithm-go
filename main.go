package main

import (
	"fmt"
)

func main() {
	s := "hello"
	for _, r := range s {
		fmt.Println(r)
	}
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
