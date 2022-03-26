package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%032b\n", int32(-2))
	fmt.Printf("%032b\n", ^int32(-2))
	fmt.Printf("%032b\n", 1 << 31)
}
