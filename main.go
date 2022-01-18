package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a uint
	var b int
	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b))

	fmt.Println(1 << 32)
}
